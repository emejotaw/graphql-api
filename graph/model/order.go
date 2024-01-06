package model

type Order struct {
	ID         string  `json:"id"`
	TotalPrice float64 `json:"totalPrice"`
}
