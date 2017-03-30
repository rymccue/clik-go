package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"strconv"
)

type Error struct {
	Err string `json:"error"`
}

func writeError(errString string, respCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(respCode) // unprocessable entity
	if err := json.NewEncoder(w).Encode(Error{errString}); err != nil {
		panic(err)
	}
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		writeError("Invalid id.", http.StatusBadRequest, w)
		return
	}

	user, err := DbGetUser(id)
	if err != nil {
		writeError("Could not find user.", http.StatusNotFound, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		writeError(err.Error(), http.StatusUnprocessableEntity, w)
		return
	}

	if err := DbCreateUser(&user); err != nil {
		writeError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func UserEdit(w http.ResponseWriter, r *http.Request) {
}

func UserGetQueue(w http.ResponseWriter, r *http.Request) {
}

func UserGetMatches(w http.ResponseWriter, r *http.Request) {
}

func DecisionCreate(w http.ResponseWriter, r *http.Request) {
}

func MatchDelete(w http.ResponseWriter, r *http.Request) {
}
