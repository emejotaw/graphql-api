package model

type Product struct {
	ID       string  `json:"id"`
	Name     *string `json:"name,omitempty"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
