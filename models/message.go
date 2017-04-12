package models

// Message represents messages table entries
type Message struct {
	ID          int64  `db:"id"`
	ChatID      int64  `db:"chat_id"`
	Message     string `db:"message"`
	MessageDate string `db:"message_date"`
}

// Messages represents a list of message objects
type Messages []Message

// MessageResponse represents a messages table entry response
type MessageResponse struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Message     string `json:"message"`
	MessageDate string `json:"message_date"`
}

// MessagesResponse represents a list of message responses
type MessagesResponse []MessageResponse
