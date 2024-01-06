package entity

import (
	"github.com/google/uuid"
)

type Product struct {
	ID       string
	Name     string
	Quantity int
	Price    float64
}

func NewProduct(name string, quantity int, price float64) *Product {

	return &Product{
		ID:       uuid.New().String(),
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}
}
