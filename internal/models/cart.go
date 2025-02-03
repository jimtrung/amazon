package models

type CartItem struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
