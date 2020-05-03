package middlewares

import (
	"errors"
	"net/http"

	"github.com/johnnyaustor/go-crud-jwt/api/auth"
	"github.com/johnnyaustor/go-crud-jwt/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		next(w, r)
	}
}
