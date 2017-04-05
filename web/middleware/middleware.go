package middleware

import (
	"net/http"

	"github.com/justinas/alice"
)

func NoMiddleware(handler http.HandlerFunc) http.Handler {
	return alice.New(NoGuestMiddleware).ThenFunc(handler)
}

func NoGuestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
