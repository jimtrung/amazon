package models

type Product struct {
	Id         string   `json:"id"`
	Image      string   `json:"image"`
	Name       string   `json:"name"`
	Rating     Rating   `json:"rating"`
	PriceCents int      `json:"priceCents"`
	Keywords   []string `json:"keywords"`
}

type Rating struct {
	Stars float64 `json:"stars"`
	Count int     `json:"count"`
}
