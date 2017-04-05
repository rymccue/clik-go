package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeffmcnd/clik/repos"
	"github.com/jeffmcnd/clik/utils"
	"github.com/jeffmcnd/clik/web/middleware"
)

func UserHandlers(r *mux.Router) {
	r.Handle("/v1/users/{id}", middleware.AuthenticatedMiddleware(GetUserHandler)).Methods("GET")
	r.Handle("/v1/users/self/edit", middleware.AuthenticatedMiddleware(EditUserHandler)).Methods("POST")
	r.Handle("/v1/users/self/queue", middleware.AuthenticatedMiddleware(GetUserQueueHandler)).Methods("GET")
	r.Handle("/v1/users/self/matches", middleware.AuthenticatedMiddleware(GetUserMatchesHandler)).Methods("GET")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	id64 := int64(id)

	if err != nil {
		utils.WriteError("Invalid id.", http.StatusBadRequest, w)
		return
	}

	user, err := repos.DbGetUser(id64)
	if err != nil {
		utils.WriteError(err.Error(), http.StatusNotFound, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
}

func GetUserQueueHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetUserIdFromRequest(r)
	if err != nil {
		utils.WriteError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	userQueue, err := repos.DbGetUserQueue(id)
	if err != nil {
		utils.WriteError(err.Error(), http.StatusNotFound, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userQueue); err != nil {
		panic(err)
	}
}

func GetUserMatchesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetUserIdFromRequest(r)
	if err != nil {
		utils.WriteError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	matches, err := repos.DbGetUserMatches(id)
	if err != nil {
		utils.WriteError(err.Error(), http.StatusNotFound, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(matches); err != nil {
		panic(err)
	}
}
