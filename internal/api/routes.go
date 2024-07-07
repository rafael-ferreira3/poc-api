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
		//Rotas desprotegidas
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", makeHTTPHandleFunc(handler.HandlerLogin))
		})

		//Rotas Protegidas
		r.Group(func(r chi.Router) {
			r.Use(AuthenticationMiddleware)
			r.Route("/user", func(r chi.Router) {
				r.Get("/", makeHTTPHandleFunc(handler.HandlerGetUsers))
				r.Get("/{id}", makeHTTPHandleFunc(handler.HandlerGetUserById))
				r.Post("/", makeHTTPHandleFunc(handler.HandlerCreateUser))
				r.Put("/", makeHTTPHandleFunc(handler.HandlerUpdate))
				r.Delete("/{id}", makeHTTPHandleFunc(handler.HandlerDeleteUser))
			})

		})
	})
}
