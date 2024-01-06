package repository

import "github.com/emejotaw/graphql-api/internal/database/entity"

type ProductRepository interface {
	Create(*entity.Product) error
	FindByID(productId string) (*entity.Product, error)
	FindAll() ([]entity.Product, error)
}
