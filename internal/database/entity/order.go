package entity

type Address struct {
	State        string
	City         string
	Neighborhood string
	Street       string
	Number       string
	ZipCode      string
}

type Order struct {
	TotalPrice float64
	Address    Address
}
