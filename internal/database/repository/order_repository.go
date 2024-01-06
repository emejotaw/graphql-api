package repository

import "github.com/emejotaw/graphql-api/internal/database/entity"

type OrderRepository interface {
	Create(order *entity.Order) error
	FindByID(orderId string) (*entity.Order, error)
	FindAll() ([]entity.Order, error)
}
