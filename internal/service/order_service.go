package service

import (
	"github.com/emejotaw/graphql-api/graph/model"
	"github.com/emejotaw/graphql-api/internal/database/entity"
	"github.com/emejotaw/graphql-api/internal/database/repository"
	"github.com/emejotaw/graphql-api/internal/database/repository/sqlite"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderService struct {
	repository repository.OrderRepository
}

func NewOrderService(db *gorm.DB) *OrderService {

	repository := sqlite.NewSqliteOrderRepository(db)
	return &OrderService{
		repository: repository,
	}
}

func (s *OrderService) Create(orderInput model.NewOrder) (*entity.Order, error) {

	order := &entity.Order{
		ID:       uuid.New().String(),
		Products: []*entity.OrderProduct{},
	}
	for _, p := range orderInput.Products {
		product := &entity.OrderProduct{
			ID:       uuid.New().String(),
			Name:     p.Name,
			Quantity: p.Quantity,
			Price:    p.Price,
		}
		order.Products = append(order.Products, product)
	}

	err := s.repository.Create(order)
	return order, err
}
