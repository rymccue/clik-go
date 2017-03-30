package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type ExtraInfo struct {
	Password string
}

func UserGet(w http.ResponseWriter, r *http.Request) {
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user User
	var info ExtraInfo
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
	if err := json.Unmarshal(body, &info); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	u := DbCreateUser(user, info.Password)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		fmt.Println(err)
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
