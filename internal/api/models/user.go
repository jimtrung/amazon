package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Country  string `json:"country"`
}

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
