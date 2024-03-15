// Package rest implements a REST API for gh-sql, mimicking the GitHub REST API.
package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gnolang/gh-sql/ent"
	"github.com/gnolang/gh-sql/ent/repository"
	"github.com/gnolang/gh-sql/ent/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Options are the server options which modify its behaviour.
type Options struct {
	DB *ent.Client
}

type server struct {
	Options
}

// Handler returns the request handler for the REST API server.
func Handler(opts Options) http.Handler {
	s := server{opts}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.handleRoot)
	r.Route("/repos/{owner}/{name}", func(r chi.Router) {
		r.Get("/", errorHandler(s.handleRepo))
		// r.Get("/issues/{issueNumber}", s.handleRepoIssues)
		// r.Get("/issues/{issueNumber}/comments", s.handleIssueComments)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/{login}", errorHandler(s.handleUser))
		// r.Get("/{login}/repos", s.handleUserRepos)
	})
	return r
}

type errCoder struct {
	code int
	error
}

func (e errCoder) Code() int { return e.code }

type writeRecorder struct {
	http.ResponseWriter
	written byte // 0 = not; 1 = written header; 2 = written header+body
}

func (wr *writeRecorder) Write(p []byte) (int, error) {
	wr.written = 2
	return wr.ResponseWriter.Write(p)
}

func (wr *writeRecorder) WriteHeader(code int) {
	if wr.written == 0 {
		wr.written = 1
	}
	wr.ResponseWriter.WriteHeader(code)
}

type errJSON struct {
	Error string `json:"error"`
}

func errorHandler(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wr := &writeRecorder{ResponseWriter: w}
		err := func() (err error) {
			defer func() {
				recErr := recover()
				if recErr != nil {
					err = fmt.Errorf("recovered: %v\n%s", err, debug.Stack())
				}
			}()
			err = fn(w, r)
			return
		}()
		if err == nil {
			return
		}
		if coder, ok := err.(interface{ Code() int }); ok {
			switch wr.written {
			case 0: // nothing written, write code + body
				w.WriteHeader(coder.Code())
				fallthrough
			case 1: // header written, write only body
				encoded, err := json.Marshal(errJSON{err.Error()})
				if err != nil {
					log.Printf("error marshaling error: %v", err)
				}
				w.Write(encoded)
			default: // just log to stderr.
				log.Printf("request error (could not write because body already written): %v", err)
			}
		} else {
			switch wr.written {
			case 0:
				w.WriteHeader(500)
				fallthrough
			case 1:
				w.Write([]byte(`{"error":"internal error"}`))
			}
			log.Printf("request error: %v", err)
		}
	}
}

func handleNotFound(err error) error {
	if ent.IsNotFound(err) {
		return errCoder{404, err}
	}
	return err
}

func (s *server) handleRoot(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("gh-sql serve\nhttps://github.com/gnolang/gh-sql\n"))
}

func encodeJSON(w http.ResponseWriter, v any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	resp, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(resp)
	return err
}

func (s *server) handleRepo(w http.ResponseWriter, r *http.Request) error {
	fullName := chi.URLParam(r, "owner") + "/" + chi.URLParam(r, "name")
	repo, err := s.DB.Repository.Query().
		Where(repository.FullName(fullName)).
		WithOwner().
		Only(r.Context())
	if err != nil {
		return handleNotFound(err)
	}

	return encodeJSON(w, repo)
}

func (s *server) handleUser(w http.ResponseWriter, r *http.Request) error {
	user, err := s.DB.User.Query().
		Where(user.Login(chi.URLParam(r, "login"))).
		Only(r.Context())
	if err != nil {
		return handleNotFound(err)
	}

	return encodeJSON(w, user)
}
