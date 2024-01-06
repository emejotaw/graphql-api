package repository

import "github.com/emejotaw/graphql-api/internal/database/entity"

type OrderProductRepository interface {
	FindByOrderID(orderID string) ([]entity.OrderProduct, error)
}
