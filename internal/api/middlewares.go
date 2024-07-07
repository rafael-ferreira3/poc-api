package api

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rafael-ferreira3/poc-api/internal/helper"
	"github.com/rafael-ferreira3/poc-api/internal/service"
)

var tokenService = service.NewTokenService()
var allowedContentTypes = []string{"application/json"}

const BEARER_PREFIX string = "Bearer "

func (s *APIServer) RegisterMiddlewares() {
	s.Router.Use(middleware.AllowContentEncoding(allowedContentTypes...))
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.WriteJson(w, http.StatusUnauthorized, ApiError{Error: "Missing Authorization Header"})
			return
		}

		if !strings.HasPrefix(authHeader, BEARER_PREFIX) {
			helper.WriteJson(w, http.StatusUnauthorized, ApiError{Error: "Invalid Authorization Header"})
			return
		}
		authHeader = authHeader[len(BEARER_PREFIX):]
		if !tokenService.VerifyToken(authHeader) {
			helper.WriteJson(w, http.StatusUnauthorized, ApiError{Error: "Invalid Authorization Token"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
