package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffmcnd/clik/models"
	"github.com/jeffmcnd/clik/repos"
	"github.com/jeffmcnd/clik/utils"
	"github.com/jeffmcnd/clik/web/middleware"
)

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
		utils.WriteError(err.Error(), http.StatusUnprocessableEntity, w)
		return
	}

	if err := repos.DbCreateUser(&user); err != nil {
		utils.WriteError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	tokenString, err := utils.GenerateTokenForUser(&user)
	if err != nil {
		utils.WriteError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	resp := models.RegisterResponse{user, tokenString}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}
