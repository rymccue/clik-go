package controllers

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

func WriteError(errString string, respCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(respCode) // unprocessable entity
	if err := json.NewEncoder(w).Encode(Result{Error: errString}); err != nil {
		panic(err)
	}
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	WriteError("Not implemented.", http.StatusInternalServerError, w)
}
