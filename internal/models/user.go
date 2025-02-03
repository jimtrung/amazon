package models

type Status int

const (
	Active Status = iota
	Inactive
	Banned
	Closed
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Country  string `json:"country"`
	Status   Status `json:"status"`
}

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
