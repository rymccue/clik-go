package models

import (
	"time"
)

type Image struct {
	UserId    int64     `json:"user_id" db:"user_id"`
	Index     int       `json:"position"`
	Small     bool      `json:"small"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	ImageUrl
}

type ImageUrl struct {
	Url string `json:"image_url"`
}

type Images []Images
