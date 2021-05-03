package middlewares

import (
	"net/http"
)

const (
	serverApiKey = "secret.secret"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientApiKey := r.Header.Get("X-API-KEY")
		if clientApiKey == "" || clientApiKey != serverApiKey {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
