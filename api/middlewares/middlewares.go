package middlewares

import (
	"DevStories/api/responses"
	"DevStories/api/tokens"
	"errors"
	"net/http"
)

//JsonMiddleware Formats all response to JSON format
func JsonMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

//AuthorizationMiddleware Checks the validity of authorization token in request header
func AuthorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tokens.IsTokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("invalid token"))
		}
		next(w, r)
	}
}
