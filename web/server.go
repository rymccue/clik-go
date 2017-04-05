package web

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeffmcnd/clik/controllers"
)

func CreateServer() *http.Server {
	r := mux.NewRouter()

	controllers.UserHandlers(r)
	controllers.DecisionHandlers(r)
	controllers.MatchHandlers(r)
	controllers.AuthHandlers(r)

	return &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
