package middleware

import (
	"net/http"

	"github.com/jeffmcnd/clik/utils"
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

func AuthenticatedMiddleware(handler http.HandlerFunc) http.Handler {
	return alice.New(AuthenticatedUserMiddleware).ThenFunc(handler)
}

func AuthenticatedUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			panic(err)
		}
		tokenString := r.Form.Get("access_token")
		if len(tokenString) == 0 {
			utils.WriteError("No access token provided.", http.StatusBadRequest, w)
			return
		}

		if err := utils.ValidateToken(tokenString); err != nil {
			utils.WriteError(err.Error(), http.StatusBadRequest, w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
