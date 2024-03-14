// Package rest implements a REST API for gh-sql, mimicking the GitHub REST API.
package rest

import (
	"net/http"
	"sync"

	"github.com/gnolang/gh-sql/ent"
)

type Options struct {
	DB *ent.Client
}

func New(opts Options) *Server {
	return &Server{
		Options: opts,
	}
}

type Server struct {
	Options

	routerOnce sync.Once
	router     *http.ServeMux
}

// ServeHTTP implements [http.Handler].
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.routerOnce.Do(func() { s.router = s.makeRouter() })
	s.router.ServeHTTP(w, r)
}

func (s *Server) makeRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleRoot)
	return mux
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("gh-sql serve\nhttps://github.com/gnolang/gh-sql\n"))
}
