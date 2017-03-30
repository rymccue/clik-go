package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func UserGet(w http.ResponseWriter, r *http.Request) {
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	if err := DbCreateUser(&user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
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
