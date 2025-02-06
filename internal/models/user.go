package models

import "time"

type User struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Country   string     `json:"country"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
