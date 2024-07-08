package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	Addr   string
	Router *chi.Mux
	server *http.Server
}

type ApiError struct {
	Error string `json:"error"`
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		Addr:   addr,
		Router: chi.NewRouter(),
	}
}

func (s *APIServer) Run() error {
	s.RegisterMiddlewares()
	s.RegisterRoutes()

	s.server = &http.Server{
		Addr:    s.Addr,
		Handler: s.Router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Listening on", s.Addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", s.Addr, err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("Server exiting")
	return nil
}
