package service

import (
	"github.com/emejotaw/graphql-api/internal/database/entity"
	"github.com/emejotaw/graphql-api/internal/database/repository"
	"github.com/emejotaw/graphql-api/internal/database/repository/sqlite"
	"gorm.io/gorm"
)

type OrderProductService struct {
	repository repository.OrderProductRepository
}

func NewOrderProductService(db *gorm.DB) *OrderProductService {

	repository := sqlite.NewSqliteOrderProductRepository(db)
	return &OrderProductService{
		repository: repository,
	}
}

func (s *OrderProductService) FindByOrderId(orderId string) ([]entity.OrderProduct, error) {

	return s.repository.FindByOrderID(orderId)
}
