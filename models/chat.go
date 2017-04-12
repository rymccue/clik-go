package models

// Chat represents an entry in the chats table
type Chat struct {
	ID      int64 `db:"id"`
	User1ID int64 `db:"user_1_id"`
	User2ID int64 `db:"user_2_id"`
}

// ChatResponse represents the JSON response for a chat
type ChatResponse struct {
	ID                int64  `json:"id"`
	UserID            int64  `json:"user_id"`
	LastMessage       string `json:"last_message"`
	LastMessageDate   string `json:"last_message_date"`
	LastMessageUserID string `json:"last_message_user_id"`
	ImageURL          string `json:"image_url"`
}

// ChatsResponse represents a list of ChatResponses
type ChatsResponse []ChatResponse
