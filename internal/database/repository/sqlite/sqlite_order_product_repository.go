package sqlite

import (
	"github.com/emejotaw/graphql-api/internal/database/entity"
	"gorm.io/gorm"
)

type SqliteOrderProductRepository struct {
	db *gorm.DB
}

func NewSqliteOrderProductRepository(db *gorm.DB) *SqliteOrderProductRepository {
	return &SqliteOrderProductRepository{
		db: db,
	}
}

func (r *SqliteOrderProductRepository) FindByOrderID(orderID string) ([]entity.OrderProduct, error) {

	products := []entity.OrderProduct{}
	err := r.db.Where("order_id = ?", orderID).Find(&products).Error
	return products, err
}
