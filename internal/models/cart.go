package models

import "time"

type Cart struct {
	CartId    int        `json:"cart_id"`
	UserId    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CartItem struct {
	CartItemId int
	CartId     int        `json:"cart_id"`
	ProductId  string     `json:"productId"`
	Quantity   int        `json:"quantity"`
	AddedAt    *time.Time `json:"added_at"`
}
