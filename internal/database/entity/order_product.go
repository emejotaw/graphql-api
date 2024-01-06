package entity

import (
	"time"
)

type OrderProduct struct {
	ID        string
	Name      string
	Quantity  int
	Price     float64
	OrderId   string
	Order     Order
	CreatedAt time.Time
}
