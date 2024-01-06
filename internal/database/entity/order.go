package entity

type Order struct {
	ID         string
	TotalPrice float64
	Products   []*OrderProduct
}
