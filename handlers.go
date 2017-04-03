package main

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	"strconv"
	"time"
)

type Error struct {
	Err string `json:"error"`
}

const (
	privateKeyPath = "keys/clik"
	publicKeyPath  = "keys/clik.pub"
)

func getKey(path string) []byte {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	info, _ := file.Stat()
	size := info.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(pembytes)

	return pembytes
}

func writeError(errString string, respCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(respCode) // unprocessable entity
	if err := json.NewEncoder(w).Encode(Error{errString}); err != nil {
		panic(err)
	}
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	tokenString := r.Form.Get("access_token")

	if len(tokenString) == 0 {
		writeError("No access token provided.", http.StatusBadRequest, w)
		return
	}

	if !ValidateToken(tokenString) {
		writeError("Invalid access token.", http.StatusBadRequest, w)
		return
	}

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
	NotImplemented(w, r)
}

func UserGetQueue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		writeError("Invalid id.", http.StatusBadRequest, w)
		return
	}

	userQueue, err := DbGetUserQueue(id)
	if err != nil {
		writeError(err.Error(), http.StatusNotFound, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userQueue); err != nil {
		panic(err)
	}
}

func UserGetMatches(w http.ResponseWriter, r *http.Request) {
	NotImplemented(w, r)
}

func DecisionCreate(w http.ResponseWriter, r *http.Request) {
	NotImplemented(w, r)
}

func MatchDelete(w http.ResponseWriter, r *http.Request) {
	NotImplemented(w, r)
}

func AccessToken(w http.ResponseWriter, r *http.Request) {
	key := getKey(privateKeyPath)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now(),
	})
	tokenString, err := token.SignedString(key)
	if err != nil {
		writeError(err.Error(), http.StatusInternalServerError, w)
	}

	if err := json.NewEncoder(w).Encode(tokenString); err != nil {
		panic(err)
	}
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	writeError("Not implemented.", http.StatusInternalServerError, w)
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return getKey(publicKeyPath), nil
	})

	if err != nil {
		return false
	}

	if token.Valid {
		return true
	}
	return false
}
