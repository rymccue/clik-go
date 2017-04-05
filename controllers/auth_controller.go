package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jeffmcnd/clik/auth"
	"github.com/jeffmcnd/clik/web/middleware"
)

func AuthHandlers(r *mux.Router) {
	r.Handle("/access-token", middleware.NoMiddleware(GetAuthTokenHandler)).Methods("GET")
}

func GetAuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	key := auth.GetKey(auth.PrivateKeyPath)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now(),
	})
	tokenString, err := token.SignedString(key)
	if err != nil {
		WriteError(err.Error(), http.StatusInternalServerError, w)
	}

	if err := json.NewEncoder(w).Encode(tokenString); err != nil {
		panic(err)
	}
}
