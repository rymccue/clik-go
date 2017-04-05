package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeffmcnd/clik/repos"
	"github.com/jeffmcnd/clik/web/middleware"
)

func MatchHandlers(r *mux.Router) {
	r.Handle("/v1/matches/{id}", middleware.NoMiddleware(DeleteMatchHandler)).Methods("DELETE")
}

func DeleteMatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		WriteError("Invalid id.", http.StatusBadRequest, w)
		return
	}

	err = repos.DbDeleteMatch(id)
	if err != nil {
		WriteError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Result{Result: "success"}); err != nil {
		panic(err)
	}
}
