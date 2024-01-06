package service

import (
	"github.com/emejotaw/graphql-api/internal/database/entity"
	"github.com/emejotaw/graphql-api/internal/database/repository"
	"github.com/emejotaw/graphql-api/internal/database/repository/sqlite"
	"gorm.io/gorm"
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(db *gorm.DB) *ProductService {

	repository := sqlite.NewSqliteProductRepository(db)
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) Create(name string, quantity int, price float64) (*entity.Product, error) {

	product := entity.NewProduct(name, quantity, price)
	err := s.repository.Create(product)
	return product, err
}

func (s *ProductService) FindByID(productId string) (*entity.Product, error) {

	return s.repository.FindByID(productId)
}

func (s *ProductService) FindAll() (*[]entity.Product, error) {

	return s.repository.FindAll()
}
