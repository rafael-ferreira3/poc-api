package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rafael-ferreira3/poc-api/internal/handler"
	"github.com/rafael-ferreira3/poc-api/internal/helper"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			helper.WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func (s *APIServer) RegisterRoutes() {

	s.Router.Route("/api/v1", func(r chi.Router) {
		r.Route("/user", func(u chi.Router) {
			u.Get("/", makeHTTPHandleFunc(handler.HandlerGetUsers))
			u.Get("/{id}", makeHTTPHandleFunc(handler.HandlerGetUserById))
			u.Post("/", makeHTTPHandleFunc(handler.HandlerCreateUser))
			u.Put("/", makeHTTPHandleFunc(handler.HandlerUpdate))
			u.Delete("/{id}", makeHTTPHandleFunc(handler.HandlerDeleteUser))
		})
	})
}
