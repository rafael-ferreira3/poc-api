package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	addr   string
	port   string
	Router *chi.Mux
}

func NewAPIServer(addr string, port string) *APIServer {
	return &APIServer{
		addr:   addr,
		port:   port,
		Router: chi.NewRouter(),
	}
}

func (s *APIServer) getUri() string {
	return s.addr + ":" + s.port
}

func (s *APIServer) Run() error {
	s.RegisterRoutes()
	log.Println("Listening on", s.getUri())
	return http.ListenAndServe(s.getUri(), s.Router)
}
