package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jeffmcnd/clik/models"
	"github.com/jeffmcnd/clik/web/middleware"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)
var upgrader = websocket.Upgrader{}

func MessageHandlers(r *mux.Router) {
	r.Handle("/v1/messages", middleware.AuthenticatedMiddleware(MessageHandler)).Methods("GET")
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true
	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
