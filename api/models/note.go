package models

import "time"

type CreateNoteRequest struct {
	UserId      int64  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type GetNoteRequest struct {
	ID          int64     `json:"first_name"`
	UserId      int64     `json:"last_name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
}
