package models

type CartItem struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
