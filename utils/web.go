package utils

import (
	"encoding/json"
	"fmt"
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

func GetTokenFromRequest(r *http.Request) (string, error) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	tokenString := r.Form.Get("access_token")
	if len(tokenString) == 0 {
		return "", fmt.Errorf("token parameter is empty or not present")
	}
	return tokenString, nil
}

func GetUserIdFromRequest(r *http.Request) (int64, error) {
	tokenString, err := GetTokenFromRequest(r)
	if err != nil {
		return -1, nil
	}
	id, err := GetUserIdFromToken(tokenString)
	if err != nil {
		return -1, nil
	}
	return id, nil
}
