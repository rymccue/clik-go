package models

type Message struct {
	UserId  int64  `json:"user_id"`
	Message string `json:"message"`
}
