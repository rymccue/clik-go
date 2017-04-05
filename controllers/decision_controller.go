package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeffmcnd/clik/models"
	"github.com/jeffmcnd/clik/repos"
	"github.com/jeffmcnd/clik/web/middleware"
)

func DecisionHandlers(r *mux.Router) {
	r.Handle("/v1/decisions/{id}", middleware.NoMiddleware(CreateDecisionHandler)).Methods("POST")
}

func CreateDecisionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		WriteError("Invalid id.", http.StatusBadRequest, w)
		return
	}

	var decisionForm models.DecisionForm
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &decisionForm); err != nil {
		WriteError(err.Error(), http.StatusUnprocessableEntity, w)
		return
	}

	err = repos.DbCreateDecision(id, &decisionForm)
	if err != nil {
		WriteError(err.Error(), http.StatusNotFound, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Result{Result: "success"}); err != nil {
		panic(err)
	}
}
