// Base HTTP clients for GitHub API

package synchub

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const (
	apiEndpoint = "https://api.github.com"
	apiVersion  = "2022-11-28"
)

// httpInternal is the internal method to call API endpoints. It sets the
// correct headers and updates the Hub's internal rate limiter.
func httpInternal(ctx context.Context, h *Hub, method, uri string, body io.Reader) (*http.Response, error) {
	// Two "flow regulators":
	// - limitsTimer ensures that we are sending requests at a sustainable pace,
	//   in accordance to GitHub's rate limit. Values are sent to it by [hub.limiter].
	// - reqsInflight ensures that we are not making too many requests to GitHub
	//   concurrently. It is only used by this function.
	h.limitsTimer <- struct{}{}
	h.reqsInflight <- struct{}{}
	defer func() {
		<-h.reqsInflight
	}()

	// set up request
	req, err := http.NewRequestWithContext(ctx, method, uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", apiVersion)
	req.Header.Set("User-Agent", "gnoverse/gh-sql")
	if h.Token != "" {
		req.Header.Set("Authorization", "Bearer "+h.Token)
	}

	if h.DebugHTTP {
		log.Printf("%s %s", req.Method, req.URL)
	}

	// do request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.Header.Get("X-Ratelimit-Limit") != "" {
		limits, err := parseLimits(resp)
		if err != nil {
			log.Printf("error parsing rate limits: %v", err)
		} else {
			h.limitsMu.Lock()
			h.limits = limits
			h.limitsMu.Unlock()
		}
	}
	return resp, nil
}

// Get performs an HTTP request to the given API endpoint on path, unmarshaling
// the result into dst.
func Get(ctx context.Context, h *Hub, path string, dst any) error {
	resp, err := httpInternal(ctx, h, "GET", apiEndpoint+path, nil)
	if err != nil {
		return err
	}

	// unmarshal body into dst
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dst)
}

func newIterWithError[T any](f func(yield func(T) bool) error) *IterWithError[T] {
	return &IterWithError[T]{i: f}
}

// IterWithError is a wrapper around an iterable function which may error.
// After iterating over [IterWithError.Values], callers should check the value
// of [IterWithError.Err].
type IterWithError[T any] struct {
	i func(yield func(T) bool) error
	// Err is set after running i.Iter.
	err error
}

// Values is an iterator function that can be yielded over.
func (i *IterWithError[T]) Values(yield func(t T) bool) {
	i.err = i.i(yield)
}

// Err returns any error that occurred while executing [IterWithError.Values].
func (i *IterWithError[T]) Err() error {
	return i.err
}

var httpLinkHeaderRe = regexp.MustCompile(`<([^>]+)>;\s*rel="([^"]+)"(?:,\s*|$)`)

// GetIterate calls the endpoint at the given path, then returns an
// [IterWithError] that can be range'd on to get all the values of the given
// resource. GetIterate automatically handles pagination.
func GetIterate[T any](ctx context.Context, h *Hub, path string) *IterWithError[T] {
	uri := apiEndpoint + path

	return newIterWithError(func(yield func(T) bool) error {
	Upper:
		for {
			resp, err := httpInternal(ctx, h, "GET", uri, nil)
			if err != nil {
				return err
			}

			// TODO(morgan): add debug flag to log failing requests.

			// retrieve data
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			_ = resp.Body.Close()

			// 100 is the max amount of elements on most GH API calls
			dst := make([]T, 0, 100)
			err = json.Unmarshal(data, &dst)
			if err != nil {
				return err
			}

			for _, item := range dst {
				if ok := yield(item); !ok {
					return nil
				}
			}

			// https://go.dev/play/p/RcjolrF-xOt
			matches := httpLinkHeaderRe.FindAllStringSubmatch(resp.Header.Get("Link"), -1)
			for _, match := range matches {
				if match[2] == "next" {
					uri = match[1]
					continue Upper
				}
			}

			// "next" link not found, return
			return nil
		}
	})
}

// -----------------------------------------------------------------------------
// Rate limiting management

type rateLimitValues struct {
	total     int
	remaining int
	reset     time.Time
}

// limiter creates a simple rate limiter based off GitHub's rate limiting.
// It divides the time up into how many requests are left in the rate limit,
// up to when it's supposed to reset.
func (h *Hub) limiter(ctx context.Context) {
	_ = h.limits // move nil check outside of loop
	for {
		h.limitsMu.Lock()
		limits := h.limits
		h.limitsMu.Unlock()

		if limits.remaining == 0 && time.Until(limits.reset) > 0 {
			dur := time.Until(limits.reset.Add(time.Second))
			log.Printf("hit rate limit, waiting until %v (%v)", limits.reset, dur)
			time.Sleep(dur)
			h.limitsMu.Lock()
			limits.remaining = limits.total
			h.limitsMu.Unlock()
			continue
		}

		if limits.remaining > starvingThreshold {
			// not starving
			// allow requests immediately to have speedy execution for
			// incremental updates.
			select {
			case <-h.limitsTimer:
			case <-ctx.Done():
				return
			}
		} else {
			// we're low on requests, perform them at a sustainable pace
			// note that Personal Access Tokens in a GH account share rate limits;
			// this allows other applications to use the same account to perform
			// operations by not instantly depleting the rate limit.
			wait := time.Until(limits.reset) / (time.Duration(limits.remaining) + 1)
			log.Printf("rate limiter is starving: waiting %s [reset: %s | remaining: %d/%d]",
				wait, limits.reset.String(), limits.remaining, limits.total)
			select {
			case <-time.After(wait):
				<-h.limitsTimer
			case <-ctx.Done():
				return
			}
		}
	}
}

func parseLimits(resp *http.Response) (vals rateLimitValues, err error) {
	h := resp.Header
	vals.total, err = strconv.Atoi(h.Get("X-Ratelimit-Limit"))
	if err != nil {
		return
	}
	vals.remaining, err = strconv.Atoi(h.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return
	}
	var reset int64
	reset, err = strconv.ParseInt(h.Get("X-Ratelimit-Reset"), 10, 64)
	if err != nil {
		return
	}
	vals.reset = time.Unix(reset, 0)
	return
}
