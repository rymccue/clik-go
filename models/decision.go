package models

type DecisionForm struct {
	UserId int64 `json:"user_id"`
	Likes  bool
}
