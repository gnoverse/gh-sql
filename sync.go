package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gnolang/gh-sql/ent"
	"github.com/peterbourgon/ff/v4"
)

func newSyncCmd(fs *ff.FlagSet, ec *execContext) *ff.Command {
	var (
		fset = ff.NewFlagSet("gh-sql sync").SetParent(fs)

		full  = fset.BoolLong("full", "sync repositories from scratch (non-incremental), automatic if more than >10_000 events to retrieve")
		token = fset.StringLong("token", "", "github personal access token to use (heavily suggested)")
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
			}.run(ctx, args)
		},
	}
}

type syncExecContext struct {
	// flags / global ctx
	*execContext
	full  bool
	token string
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

		hub.fetchRepository(ctx, parts[0], parts[1])
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

func (s *syncHub) get(ctx context.Context, path string, dst any) error {
	// block until the rate limiter allows us to do request
	s.requestBucket <- struct{}{}

	// set up request
	req, err := http.NewRequestWithContext(ctx, "GET", apiEndpoint+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", apiVersion)
	req.Header.Set("User-Agent", "gnolang/gh-sql")
	if s.token != "" {
		req.Header.Set("Authorization", "Bearer "+s.token)
	}

	// do request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
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

	// unmarshal body into dst
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dst)
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

func (s *syncHub) fetchRepository(ctx context.Context, owner, repo string) {
	// determine if we should fetch the repository, or it's already been updated
	keyVal := fmt.Sprintf("/repos/%s/%s", owner, repo)
	if !s.shouldFetch(keyVal) {
		return
	}

	// wait to run goroutine, and execute.
	s.runWait()
	go s._fetchRepository(ctx, owner, repo)
}

func (s *syncHub) _fetchRepository(ctx context.Context, owner, repo string) {
	defer s.recover()

	var r ent.Repository
	if err := s.get(ctx, fmt.Sprintf("/repos/%s/%s", owner, repo), &r); err != nil {
		s.report(fmt.Errorf("fetchRepository(%q, %q) get: %w", owner, repo, err))
		return
	}
	err := s.db.Repository.Create().
		CopyRepository(&r).
		OnConflict().UpdateNewValues().
		Exec(ctx)
	if err != nil {
		s.report(fmt.Errorf("fetchRepository(%q, %q) save: %w", owner, repo, err))
		return
	}
}
