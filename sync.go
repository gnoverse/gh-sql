package main

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
	"github.com/peterbourgon/ff/v4"
)

func newSyncCmd(fs *ff.FlagSet, ec *execContext) *ff.Command {
	var (
		fset = ff.NewFlagSet("gh-sql sync").SetParent(fs)

		full  = fset.BoolLong("full", "sync repositories from scratch (non-incremental), automatic if more than >10_000 events to retrieve")
		token = fset.StringLong("token", "", "github personal access token to use (heavily suggested)")

		debugHTTP = fset.BoolLong("debug.http", "log http requests")
	)
	return &ff.Command{
		Name:      "sync",
		Usage:     "gh-sql sync [FLAGS...] REPOS...",
		ShortHelp: "synchronize the current state with GitHub",
		LongHelp:  "Repositories must be specified with the syntax <owner>/<repo>.",
		Flags:     fset,
		Exec: func(ctx context.Context, args []string) error {
			return syncExecContext{
				ec,
				*full,
				*token,
				*debugHTTP,
			}.run(ctx, args)
		},
	}
}

type syncExecContext struct {
	// flags / global ctx
	*execContext
	full      bool
	token     string
	debugHTTP bool
}

func (s syncExecContext) run(ctx context.Context, args []string) error {
	values := rateLimitValues{
		// GitHub defaults for unauthenticated
		total:     60,
		remaining: 60,
		reset:     time.Now().Truncate(time.Hour).Add(time.Hour),
	}
	if s.token != "" {
		// GitHub defaults for authenticated
		values.total = 5000
		values.remaining = 5000
	}

	// set up hub
	hub := &syncHub{
		syncExecContext: s,
		running:         make(chan struct{}, 8),
		requestBucket:   make(chan struct{}, 8),
		limits:          values,
		updated:         make(map[string]struct{}),
	}
	go hub.limiter(ctx)

	// execute
	for _, arg := range args {
		parts := strings.SplitN(arg, "/", 2)
		if len(parts) != 2 {
			log.Printf("invalid repo syntax: %q (must be <owner>/<repo>)", arg)
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

type syncHub struct {
	syncExecContext

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
func (s *syncHub) limiter(ctx context.Context) {
	_ = s.limits // move nil check outside of loop
	for {
		s.limitsMu.Lock()
		limits := s.limits
		s.limitsMu.Unlock()

		if limits.remaining == 0 && time.Until(limits.reset) > 0 {
			dur := time.Until(limits.reset)
			log.Printf("hit rate limit, waiting until %v (%v)", limits.reset, dur)
			time.Sleep(dur)
			s.limitsMu.Lock()
			limits.remaining = limits.total
			s.limitsMu.Unlock()
			continue
		}

		wait := time.Until(limits.reset) / (time.Duration(limits.remaining) + 1)
		wait = max(wait, 100*time.Millisecond)
		select {
		case <-time.After(wait):
			<-s.requestBucket
		case <-ctx.Done():
			return
		}
	}
}

// runWait waits until a spot frees up in s.running, useful before
// starting a new goroutine.
func (s *syncHub) runWait() {
	s.running <- struct{}{}
	s.wg.Add(1)
}

// recover performs error recovery in s.fetch* functions,
// as well as freeing up a space for running goroutines.
// it should be ran as a deferred function in all fetch*
// functions.
func (s *syncHub) recover() {
	if err := recover(); err != nil {
		trace := debug.Stack()
		log.Printf("panic: %v\n%v", err, string(trace))
	}
	s.wg.Done()
	<-s.running
}

func (s *syncHub) shouldFetch(key string) (should bool) {
	s.updatedMu.Lock()
	_, ok := s.updated[key]
	if !ok {
		s.updated[key] = struct{}{}
		should = true
	}
	s.updatedMu.Unlock()
	return
}

// report adds an error to s
func (s *syncHub) report(err error) {
	log.Println("error:", err)
}

const (
	apiEndpoint = "https://api.github.com"
	apiVersion  = "2022-11-28"
)

func httpGet(ctx context.Context, s *syncHub, path string, dst any) error {
	resp, err := httpInternal(ctx, s, "GET", apiEndpoint+path, nil)

	// unmarshal body into dst
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dst)
}

func httpInternal(ctx context.Context, s *syncHub, method, uri string, body io.Reader) (*http.Response, error) {
	// block until the rate limiter allows us to do request
	s.requestBucket <- struct{}{}

	// set up request
	req, err := http.NewRequestWithContext(ctx, method, uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", apiVersion)
	req.Header.Set("User-Agent", "gnolang/gh-sql")
	if s.token != "" {
		req.Header.Set("Authorization", "Bearer "+s.token)
	}

	if s.debugHTTP {
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
			s.limitsMu.Lock()
			s.limits = limits
			s.limitsMu.Unlock()
		}
	}
	return resp, nil
}

var httpLinkHeaderRe = regexp.MustCompile(`<([^>]+)>;\s*rel="([^"]+)"(?:,\s*|$)`)

func httpGetIterate[T any](ctx context.Context, s *syncHub, path string, fn func(item T) error) error {
	uri := apiEndpoint + path
	for {
		resp, err := httpInternal(ctx, s, "GET", uri, nil)

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
		// for database errors and it avoids us having to pass `*syncHub` to this fn.
		panic(fmt.Errorf("runHooks(%+v) query: %w", key, err))
	}
	f(result)
}

func (s *syncHub) fetchRepository(ctx context.Context, owner, repo string) {
	// determine if we should fetch the repository, or it's already been updated
	keyVal := fmt.Sprintf("/repos/%s/%s", owner, repo)
	if !s.shouldFetch(keyVal) {
		runHooks(ctx,
			hookRepository{owner: owner, repo: repo},
			s.db.Repository.Query().
				Where(repository.FullName(owner+"/"+repo)).Only,
		)
		return
	}

	// wait to run goroutine, and execute.
	s.runWait()
	go s._fetchRepository(ctx, owner, repo)
}

func (s *syncHub) _fetchRepository(ctx context.Context, owner, repo string) {
	defer s.recover()

	var r struct {
		ent.Repository
		Owner struct {
			Login string `json:"login"`
		} `json:"owner"`
	}
	if err := httpGet(ctx, s, fmt.Sprintf("/repos/%s/%s", owner, repo), &r); err != nil {
		s.report(fmt.Errorf("fetchRepository(%q, %q) get: %w", owner, repo, err))
		return
	}
	err := s.db.Repository.Create().
		CopyRepository(&r.Repository).
		OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		s.report(fmt.Errorf("fetchRepository(%q, %q) save: %w", owner, repo, err))
		return
	}

	runHooks(ctx,
		hookRepository{owner: owner, repo: repo},
		s.db.Repository.Query().
			Where(repository.FullName(owner+"/"+repo)).Only,
	)

	s.fetchUser(
		addHook(ctx, hookUser{username: r.Owner.Login}, func(u *ent.User) {
			err := s.db.Repository.Update().
				Where(repository.ID(r.ID)).
				SetOwnerID(u.ID).
				Exec(ctx)
			if err != nil {
				s.report(fmt.Errorf("link repo %d (%s) to owner %d (%s): %w",
					r.ID, r.Name, u.ID, u.Login, err))
			}
		}),
		r.Owner.Login,
	)
}

func (s *syncHub) fetchRepositories(ctx context.Context, owner string, shouldStore func(*ent.Repository) bool) {
	// wait to run goroutine, and execute.
	s.runWait()
	go s._fetchRepositories(ctx, owner, shouldStore)
}

func (s *syncHub) _fetchRepositories(ctx context.Context, owner string, shouldFetch func(*ent.Repository) bool) {
	defer s.recover()

	err := httpGetIterate(ctx, s, fmt.Sprintf("/users/%s/repos?per_page=100", owner), func(r *ent.Repository) error {
		if shouldFetch(r) {
			s.fetchRepository(ctx, owner, r.Name)
		}
		return nil
	})
	if err != nil {
		s.report(fmt.Errorf("fetchRepositories(%q): %w", owner, err))
	}
}

func (s *syncHub) fetchUser(ctx context.Context, username string) {
	// determine if we should fetch the user, or it's already been updated
	keyVal := fmt.Sprintf("/users/%s", username)
	if !s.shouldFetch(keyVal) {
		runHooks(ctx,
			hookUser{username: username},
			s.db.User.Query().
				Where(user.Login(username)).Only,
		)
		return
	}

	// wait to run goroutine, and execute.
	s.runWait()
	go s._fetchUser(ctx, username)
}

func (s *syncHub) _fetchUser(ctx context.Context, username string) {
	defer s.recover()

	var u ent.User
	if err := httpGet(ctx, s, fmt.Sprintf("/users/%s", username), &u); err != nil {
		s.report(fmt.Errorf("fetchUser(%q) get: %w", username, err))
		return
	}
	err := s.db.User.Create().
		CopyUser(&u).
		OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		s.report(fmt.Errorf("fetchUser(%q) save: %w", username, err))
		return
	}
	runHooks(ctx,
		hookUser{username: username},
		s.db.User.Query().
			Where(user.Login(username)).Only,
	)
}
