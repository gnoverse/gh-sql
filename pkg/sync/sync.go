// Package sync implements a system to fetch GitHub repositories concurrently.
package sync

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gnolang/gh-sql/ent"
	"github.com/gnolang/gh-sql/ent/repository"
	"github.com/gnolang/gh-sql/ent/user"
)

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

	// set up hub
	hub := &hub{
		Options:       opts,
		running:       make(chan struct{}, 8),
		requestBucket: make(chan struct{}, 8),
		limits:        values,
		updated:       make(map[string]struct{}),
	}
	go hub.limiter(ctx)

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
			hub.fetchRepositories(ctx, parts[0], func(r *ent.Repository) bool {
				return re.MatchString(r.Name)
			})
		} else {
			// no wildcards; fetch repo directly.
			hub.fetchRepository(ctx, parts[0], parts[1])
		}
	}

	// wait for all goros to finish
	hub.wg.Wait()

	return nil
}

// Options repsent the options to run Sync. These match the flags provided on
// the command line.
type Options struct {
	DB    *ent.Client
	Full  bool
	Token string

	DebugHTTP bool
}

type hub struct {
	Options

	// synchronization/limiting channels
	// number of running goroutines
	running chan struct{}
	wg      sync.WaitGroup
	// channel used to space out requests to the GitHub API
	requestBucket chan struct{}

	// rate limits in place
	limits rateLimitValues
	// only at most 1 reader, so RWMutex isn't needed
	limitsMu sync.Mutex

	// list of updated resources, de-duplicating updates.
	updated   map[string]struct{}
	updatedMu sync.Mutex
}

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
		wait = max(wait, 100*time.Millisecond)
		select {
		case <-time.After(wait):
			<-h.requestBucket
		case <-ctx.Done():
			return
		}
	}
}

// runWait waits until a spot frees up in s.running, useful before
// starting a new goroutine.
func (h *hub) runWait() {
	h.running <- struct{}{}
	h.wg.Add(1)
}

// recover performs error recovery in s.fetch* functions,
// as well as freeing up a space for running goroutines.
// it should be ran as a deferred function in all fetch*
// functions.
func (h *hub) recover() {
	if err := recover(); err != nil {
		trace := debug.Stack()
		log.Printf("panic: %v\n%v", err, string(trace))
	}
	h.wg.Done()
	<-h.running
}

func (h *hub) shouldFetch(key string) (should bool) {
	h.updatedMu.Lock()
	_, ok := h.updated[key]
	if !ok {
		h.updated[key] = struct{}{}
		should = true
	}
	h.updatedMu.Unlock()
	return
}

// report adds an error to s
func (h *hub) report(err error) {
	log.Println("error:", err)
}

const (
	apiEndpoint = "https://api.github.com"
	apiVersion  = "2022-11-28"
)

func httpGet(ctx context.Context, h *hub, path string, dst any) error {
	resp, err := httpInternal(ctx, h, "GET", apiEndpoint+path, nil)

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
	for {
		resp, err := httpInternal(ctx, h, "GET", uri, nil)

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
				continue
			}
		}

		// "next" link not found, return
		return nil
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
	reset, err = strconv.ParseInt(h.Get("X-Ratelimit-Remaining"), 10, 64)
	if err != nil {
		return
	}
	vals.reset = time.Unix(reset, 0)
	return
}

// hookKey is an interface implemented by hook* types, which can be used as
// context.Context values.
//
// hookKey enables to type check, through generics, the associated functions
// and keys passed to addHook and runHooks.
type hookKey[T any] interface{ hookKey(T) }

type hookImpl[T any] struct{}

func (hookImpl[T]) hookKey(T) {}

type (
	// hookRepository is called after a repository has been successfully
	// fetched.
	hookRepository struct {
		hookImpl[*ent.Repository]
		owner, repo string
	}
	// hookUser is called after a user has been successfully fetched.
	hookUser struct {
		hookImpl[*ent.User]
		username string
	}
)

// addHook adds a new hook for the "event" at key, which on success will
// execute the function f.
func addHook[T any](ctx context.Context, key hookKey[T], f func(T)) context.Context {
	v := ctx.Value(key)
	if v != nil {
		f2, ok := v.(func(T))
		if !ok {
			panic(fmt.Errorf("addHook(%+v): invalid function for %T hook: %T", key, key, f))
		}
		return context.WithValue(ctx, key, func(val T) {
			f(val)
			f2(val)
		})
	}
	return context.WithValue(ctx, key, f)
}

// runHooks calls the hooks associated with the given key.
// query should be a closure to an ent function call, to retrieve the given
// value. It is only executed if there are hooked functions.
func runHooks[T any](ctx context.Context, key hookKey[T], query func(ctx context.Context) (T, error)) {
	v := ctx.Value(key)
	if v == nil {
		return
	}
	f, ok := v.(func(T))
	if !ok {
		panic(fmt.Errorf("runHooks(%+v): invalid function for %T hook: %T", key, key, f))
	}
	result, err := query(ctx)
	if err != nil {
		// This should be an error rather than a panic, but normally this should happen only
		// for database errors and it avoids us having to pass `*hub` to this fn.
		panic(fmt.Errorf("runHooks(%+v) query: %w", key, err))
	}
	f(result)
}

func (h *hub) fetchRepository(ctx context.Context, owner, repo string) {
	// determine if we should fetch the repository, or it's already been updated
	keyVal := fmt.Sprintf("/repos/%s/%s", owner, repo)
	if !h.shouldFetch(keyVal) {
		runHooks(ctx,
			hookRepository{owner: owner, repo: repo},
			h.DB.Repository.Query().
				Where(repository.FullName(owner+"/"+repo)).Only,
		)
		return
	}

	// wait to run goroutine, and execute.
	h.runWait()
	go h._fetchRepository(ctx, owner, repo)
}

func (h *hub) _fetchRepository(ctx context.Context, owner, repo string) {
	defer h.recover()

	var r struct {
		ent.Repository
		Owner struct {
			Login string `json:"login"`
		} `json:"owner"`
	}
	if err := httpGet(ctx, h, fmt.Sprintf("/repos/%s/%s", owner, repo), &r); err != nil {
		h.report(fmt.Errorf("fetchRepository(%q, %q) get: %w", owner, repo, err))
		return
	}
	err := h.DB.Repository.Create().
		CopyRepository(&r.Repository).
		OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		h.report(fmt.Errorf("fetchRepository(%q, %q) save: %w", owner, repo, err))
		return
	}

	runHooks(ctx,
		hookRepository{owner: owner, repo: repo},
		h.DB.Repository.Query().
			Where(repository.FullName(owner+"/"+repo)).Only,
	)

	h.fetchUser(
		addHook(ctx, hookUser{username: r.Owner.Login}, func(u *ent.User) {
			err := h.DB.Repository.Update().
				Where(repository.ID(r.ID)).
				SetOwnerID(u.ID).
				Exec(ctx)
			if err != nil {
				h.report(fmt.Errorf("link repo %d (%s) to owner %d (%s): %w",
					r.ID, r.Name, u.ID, u.Login, err))
			}
		}),
		r.Owner.Login,
	)
}

func (h *hub) fetchRepositories(ctx context.Context, owner string, shouldStore func(*ent.Repository) bool) {
	// wait to run goroutine, and execute.
	h.runWait()
	go h._fetchRepositories(ctx, owner, shouldStore)
}

func (h *hub) _fetchRepositories(ctx context.Context, owner string, shouldFetch func(*ent.Repository) bool) {
	defer h.recover()

	err := httpGetIterate(ctx, h, fmt.Sprintf("/users/%s/repos?per_page=100", owner), func(r *ent.Repository) error {
		if shouldFetch(r) {
			h.fetchRepository(ctx, owner, r.Name)
		}
		return nil
	})
	if err != nil {
		h.report(fmt.Errorf("fetchRepositories(%q): %w", owner, err))
	}
}

func (h *hub) fetchUser(ctx context.Context, username string) {
	// determine if we should fetch the user, or it's already been updated
	keyVal := fmt.Sprintf("/users/%s", username)
	if !h.shouldFetch(keyVal) {
		runHooks(ctx,
			hookUser{username: username},
			h.DB.User.Query().
				Where(user.Login(username)).Only,
		)
		return
	}

	// wait to run goroutine, and execute.
	h.runWait()
	go h._fetchUser(ctx, username)
}

func (h *hub) _fetchUser(ctx context.Context, username string) {
	defer h.recover()

	var u ent.User
	if err := httpGet(ctx, h, fmt.Sprintf("/users/%s", username), &u); err != nil {
		h.report(fmt.Errorf("fetchUser(%q) get: %w", username, err))
		return
	}
	err := h.DB.User.Create().
		CopyUser(&u).
		OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		h.report(fmt.Errorf("fetchUser(%q) save: %w", username, err))
		return
	}
	runHooks(ctx,
		hookUser{username: username},
		h.DB.User.Query().
			Where(user.Login(username)).Only,
	)
}
