package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jeffmcnd/clik/auth"
	"github.com/jeffmcnd/clik/models"
	"github.com/jeffmcnd/clik/repos"
	"github.com/jeffmcnd/clik/web/middleware"
)

type RegisterResponse struct {
	AccessToken string `json:"access_token"`
}

func AuthHandlers(r *mux.Router) {
	r.Handle("/register", middleware.NoMiddleware(RegisterHandler)).Methods("POST")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		WriteError(err.Error(), http.StatusUnprocessableEntity, w)
		return
	}

	if err := repos.DbCreateUser(&user); err != nil {
		WriteError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	key := auth.GetKey(auth.PrivateKeyPath)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":   time.Now(),
		"exp":   time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
		"email": user.Email,
	})
	tokenString, err := token.SignedString(key)
	if err != nil {
		WriteError(err.Error(), http.StatusInternalServerError, w)
	}

	resp := RegisterResponse{tokenString}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}
