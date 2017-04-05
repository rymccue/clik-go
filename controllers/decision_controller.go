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

func DecisionHandlers(r *mux.Router) {
	r.Handle("/v1/decisions", middleware.AuthenticatedMiddleware(CreateDecisionHandler)).Methods("POST")
}

func CreateDecisionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetUserIdFromRequest(r)
	if err != nil {
		utils.WriteError("Invalid id.", http.StatusBadRequest, w)
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
		utils.WriteError(err.Error(), http.StatusUnprocessableEntity, w)
		return
	}

	err = repos.DbCreateDecision(id, &decisionForm)
	if err != nil {
		utils.WriteError(err.Error(), http.StatusNotFound, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(utils.Result{Result: "success"}); err != nil {
		panic(err)
	}
}
