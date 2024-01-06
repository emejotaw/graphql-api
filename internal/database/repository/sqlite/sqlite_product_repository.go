package sqlite

import (
	"github.com/emejotaw/graphql-api/internal/database/entity"
	"gorm.io/gorm"
)

type SqliteProductRepository struct {
	db *gorm.DB
}

func NewSqliteProductRepository(db *gorm.DB) *SqliteProductRepository {
	return &SqliteProductRepository{
		db: db,
	}
}

func (r *SqliteProductRepository) Create(product *entity.Product) error {

	return r.db.Create(product).Error
}

func (r *SqliteProductRepository) FindByID(productId string) (*entity.Product, error) {

	product := &entity.Product{}
	err := r.db.First(product, productId).Error
	return product, err
}

func (r *SqliteProductRepository) FindAll() (*[]entity.Product, error) {

	products := &[]entity.Product{}
	err := r.db.Find(products).Error

	return products, err
}
