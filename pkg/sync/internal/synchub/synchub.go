// Package synchub provides a data structure to manage:
//
//   - Rate limiting of requests to the GitHub API.
//   - Memoization of request results.
//   - Management of goroutines and inflight actions, through an errgroup.
package synchub

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gnoverse/gh-sql/ent"
	"golang.org/x/sync/errgroup"
)

const (
	maxConcurrentRequests = 4

	// maximum requests per hour when unauthenticated.
	unauthenticatedMaxRequests = 60
	// maximum requests per hour when authenticated.
	authenticatedMaxRequests = 5000

	// When the number of requests left drops below the threshold, then we are
	// "starving", and the rate limiter properly kicks in.
	starvingThreshold = 2000
)

// Options repsent the options for creating a hub.
type Options struct {
	DB    *ent.Client // TODO: find a way to move it out of this package.
	Token string

	DebugHTTP bool
}

// New creates a new synchub, with the given options.
func New(ctx context.Context, opts Options) *Hub {
	values := rateLimitValues{
		// GitHub defaults for unauthenticated
		total:     unauthenticatedMaxRequests,
		remaining: unauthenticatedMaxRequests,
		reset:     time.Now().Truncate(time.Minute).Add(time.Hour),
	}
	if opts.Token != "" {
		// GitHub defaults for authenticated
		values.total = authenticatedMaxRequests
		values.remaining = authenticatedMaxRequests
	}

	gr, ctx := errgroup.WithContext(ctx)

	// set up hub
	h := &Hub{
		Options:      opts,
		Group:        gr,
		limits:       values,
		limitsTimer:  make(chan struct{}, 1),
		reqsInflight: make(chan struct{}, maxConcurrentRequests),
		updated:      make(map[string]any),
	}
	go h.limiter(ctx)

	return h
}

// Hub is the main data structure provided by this package, to manage rate
// limiting and request memoization.
type Hub struct {
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
	// limitsTimer regulates how requests can be sent to the GitHub API.
	limitsTimer chan struct{}
	// reqsInflight keeps track of how many inflights requests there are,
	// and limits them to the chan's cap.
	reqsInflight chan struct{}

	// list of updated resources, de-duplicating updates.
	// Keys are the values of ID() of each [resource] being fetched.
	// Values are of type func() (T, error)
	updated   map[string]any
	updatedMu sync.Mutex
}

// Resource represents the base API interface of a resource.
type Resource[T any] interface {
	// ID should be a unique path identifying this resource, used to avoid
	// duplication of requests in a single Sync run. This can generally simply
	// be the path.
	ID() string
	// Fetch retrieves the resource from the HTTP API.
	Fetch(ctx context.Context, h *Hub) (T, error)
}

// Fetch calls Fetch on res if there is no memoized resource with the same
// [Resource.ID]; otherwise, itreturns a previously memoized value from the [Hub] h.
func Fetch[T any](ctx context.Context, h *Hub, res Resource[T]) (T, error) {
	fn, _ := SetGetter(h, res.ID(), func() (T, error) {
		return res.Fetch(ctx, h)
	})

	return fn()
}

// FetchAsync calls Fetch on the given resource res if necessary. It performs
// this in another goroutine. No goroutine is created if Fetch or FetchAsync
// were already called on a resource with the same [Resource.ID].
func FetchAsync[T any](ctx context.Context, h *Hub, res Resource[T]) {
	fn, created := SetGetter(h, res.ID(), func() (T, error) {
		return res.Fetch(ctx, h)
	})

	if created {
		h.Go(func() error {
			_, err := fn()
			// TODO: configure (halt at first error or not)
			if err != nil {
				h.Warn(err)
			}
			return nil
		})
	}
}

// SetGetter checks whether the internal map of updated resources already
// contains a value with the given id. If it doesn't exist, fn is used as the
// getter; otherwise the existing value is returned. The values are internally
// wrapped in [sync.OnceValues], meaning that the values returned by the getter
// a first time are "memoized"; so retrieval doesn't happen more than once.
//
// SetGetter is generally used internally by [Fetch] and [FetchAsync];
// it should be used elsewhere when the full resource in question has been
// retrieved otherwise, and Fetch/FetchAsync is not necessary.
func SetGetter[T any](h *Hub, id string, fn func() (T, error)) (getter func() (T, error), created bool) {
	h.updatedMu.Lock()
	getterRaw, ok := h.updated[id]
	if !ok {
		getter = sync.OnceValues(fn)
		h.updated[id] = getter
		created = true
	} else {
		getter = getterRaw.(func() (T, error))
	}
	h.updatedMu.Unlock()
	return
}

// Warn prints the given error as a warning, without halting execution.
func (h *Hub) Warn(err error) {
	log.Println("warning:", err)
}
