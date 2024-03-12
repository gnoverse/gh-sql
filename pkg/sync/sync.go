// Package sync implements a system to fetch GitHub repositories concurrently.
package sync

// This file contains primarily the "infrastructure" for the sync system
// (central functions for http requests, concurrency management, etc.)
// and the main exported entrypoint (Sync).

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gnolang/gh-sql/ent"
	"golang.org/x/sync/errgroup"
)

// Options repsent the options to run Sync. These match the flags provided on
// the command line.
type Options struct {
	DB    *ent.Client
	Full  bool
	Token string

	DebugHTTP bool
}

// Sync performs the synchronisation of repositories to the database provided
// in the options.
func Sync(ctx context.Context, repositories []string, opts Options) error {
	values := rateLimitValues{
		// GitHub defaults for unauthenticated
		total:     60,
		remaining: 60,
		reset:     time.Now().Truncate(time.Hour).Add(time.Hour),
	}
	if opts.Token != "" {
		// GitHub defaults for authenticated
		values.total = 5000
		values.remaining = 5000
	}

	gr, ctx := errgroup.WithContext(ctx)

	// set up hub
	h := &hub{
		Options:       opts,
		Group:         gr,
		requestBucket: make(chan struct{}, 8),
		limits:        values,
		updated:       make(map[string]any),
	}
	go h.limiter(ctx)

	// execute
	for _, repo := range repositories {
		parts := strings.SplitN(repo, "/", 2)
		if len(parts) != 2 {
			log.Printf("invalid repo syntax: %q (must be <owner>/<repo>)", repo)
		}

		if strings.IndexByte(parts[1], '*') != -1 {
			// contains wildcards
			reString := "^" + strings.ReplaceAll(regexp.QuoteMeta(parts[1]), `\*`, `.*`) + "$"
			re := regexp.MustCompile(reString)
			fetchRepositories(ctx, h, parts[0], func(r *ent.Repository) bool {
				return re.MatchString(r.Name)
			})
		} else {
			// no wildcards; fetch repo directly.
			fetchAsync(ctx, h, fetchRepository{owner: parts[0], repo: parts[1]})
		}
	}

	// wait for all goros to finish
	return h.Wait()
}

// -----------------------------------------------------------------------------
// Hub and concurrency helpers

type hub struct {
	Options

	// errgroup, for launching goroutines.
	// NOTE: the errgroup should be used when launching goroutines to fetch
	// indefinite numbers of resources, for instance, when iterating over a
	// list returned by the GitHub API. If a resource needs to fetch its own
	// fixed dependent resource, it may do so independently without using this
	// Group to limit its goroutines.
	*errgroup.Group

	// rate limits in place
	limits rateLimitValues
	// only at most 1 reader, so RWMutex isn't needed
	limitsMu sync.Mutex
	// channel used to space out requests to the GitHub API
	requestBucket chan struct{}

	// list of updated resources, de-duplicating updates.
	// Keys are the values of ID() of each [resource] being fetched.
	// Values are of type func() (T, error)
	updated   map[string]any
	updatedMu sync.Mutex
}

// resource represents the base API interface of a resource.
type resource[T any] interface {
	// ID should be a unique path identifying this resource, used to avoid
	// duplication of requests in a single Sync run. This can generally simply
	// be the path.
	ID() string
	// Fetch retrieves the resource from the HTTP API.
	Fetch(ctx context.Context, h *hub) (T, error)
}

func fetch[T any](ctx context.Context, h *hub, res resource[T]) (T, error) {
	fn, _ := fetchInternal(ctx, h, res)

	return fn()
}

func fetchAsync[T any](ctx context.Context, h *hub, res resource[T]) {
	fn, created := fetchInternal(ctx, h, res)

	if created {
		h.Go(func() error {
			_, err := fn()
			// TODO: configure (halt at first error or not)
			if err != nil {
				h.warn(err)
			}
			return nil
		})
	}
}

func fetchInternal[T any](ctx context.Context, h *hub, res resource[T]) (func() (T, error), bool) {
	id := res.ID()

	var created bool
	h.updatedMu.Lock()
	fn, ok := h.updated[id]
	if !ok {
		fn = sync.OnceValues(func() (T, error) {
			return res.Fetch(ctx, h)
		})
		h.updated[id] = fn
		created = true
	}
	h.updatedMu.Unlock()

	return fn.(func() (T, error)), created
}

// warn prints the given error as a warning, without halting execution.
func (h *hub) warn(err error) {
	log.Println("warning:", err)
}

// -----------------------------------------------------------------------------
// Base HTTP clients for GitHub API

const (
	apiEndpoint = "https://api.github.com"
	apiVersion  = "2022-11-28"
)

func httpGet(ctx context.Context, h *hub, path string, dst any) error {
	resp, err := httpInternal(ctx, h, "GET", apiEndpoint+path, nil)
	if err != nil {
		return err
	}

	// unmarshal body into dst
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dst)
}

func httpInternal(ctx context.Context, h *hub, method, uri string, body io.Reader) (*http.Response, error) {
	// block until the rate limiter allows us to do request
	h.requestBucket <- struct{}{}

	// set up request
	req, err := http.NewRequestWithContext(ctx, method, uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", apiVersion)
	req.Header.Set("User-Agent", "gnolang/gh-sql")
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

var httpLinkHeaderRe = regexp.MustCompile(`<([^>]+)>;\s*rel="([^"]+)"(?:,\s*|$)`)

func httpGetIterate[T any](ctx context.Context, h *hub, path string, fn func(item T) error) error {
	uri := apiEndpoint + path
Upper:
	for {
		resp, err := httpInternal(ctx, h, "GET", uri, nil)
		if err != nil {
			return err
		}

		// retrieve data
		data, err := ioutil.ReadAll(resp.Body)
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
			if err := fn(item); err != nil {
				return err
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
func (h *hub) limiter(ctx context.Context) {
	_ = h.limits // move nil check outside of loop
	for {
		h.limitsMu.Lock()
		limits := h.limits
		h.limitsMu.Unlock()

		if limits.remaining == 0 && time.Until(limits.reset) > 0 {
			dur := time.Until(limits.reset)
			log.Printf("hit rate limit, waiting until %v (%v)", limits.reset, dur)
			time.Sleep(dur)
			h.limitsMu.Lock()
			limits.remaining = limits.total
			h.limitsMu.Unlock()
			continue
		}

		wait := time.Until(limits.reset) / (time.Duration(limits.remaining) + 1)
		if h.DebugHTTP {
			log.Printf("rate limit delay: %v (resets %s)", wait, limits.reset.String())
		}
		wait = max(wait, 100*time.Millisecond)
		select {
		case <-time.After(wait):
			<-h.requestBucket
		case <-ctx.Done():
			return
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
