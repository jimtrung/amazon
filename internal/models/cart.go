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
	ProductId  string     `json:"product_id"`
	Quantity   int        `json:"quantity"`
	AddedAt    *time.Time `json:"added_at"`
}

type CartWithItems struct {
	CartItemId    int       `json:"cart_item_id"`
	CartId        int       `json:"cart_id"`
	ProductId     string    `json:"product_id"`
	Quantity      int       `json:"quantity"`
	AddedAt       time.Time `json:"added_at"`
	UserId        int       `json:"user_id"`
	CartCreatedAt time.Time `json:"cart_created_at"`
	CartUpdatedAt time.Time `json:"cart_updated_at"`
}
