package models

import "time"

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateUser model is only used when create a user.
// It has validations defined.
type CreateUser struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Username  string    `json:"username" validate:"required,min=5,max=20"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8,containsany=!@#$%^&*(),upper,lower,number"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	Username     string `json:"username"`
	ID           int    `json:"id"`
	AccessToken  string `json:"access_token"`
	RefershToken string `json:"refersh_token"`
}
