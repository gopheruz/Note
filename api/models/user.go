package models

import "time"

type CreateUserRequest struct {
	FirstName   string  `json:"first_name" binding:"required,min=2,max=50"`
	LastName    string  `json:"last_name" binding:"required,min=2,max=50"`
	Email       string  `json:"email" binding:"required,email,min=15,max=100"`
	Password    string  `json:"password" binding:"required,min=6,max=16"`
	PhoneNumber *string `json:"phone_number" binding:"min=9,max=20"`
	ImageUrl    *string `json:"image_url"`
}

type GetUserResponse struct {
	ID          int64      `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	PhoneNumber *string    `json:"phone_number"`
	ImageUrl    *string    `json:"image_url"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
