package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	addr   string
	Router *chi.Mux
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr:   addr,
		Router: chi.NewRouter(),
	}
}

func (s *APIServer) Run() error {
	s.RegisterRoutes()
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, s.Router)
}
