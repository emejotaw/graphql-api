package sqlite

import (
	"github.com/emejotaw/graphql-api/internal/database/entity"
	"gorm.io/gorm"
)

type SqliteOrderRepository struct {
	db *gorm.DB
}

func NewSqliteOrderRepository(db *gorm.DB) *SqliteOrderRepository {

	return &SqliteOrderRepository{
		db: db,
	}
}

func (r *SqliteOrderRepository) Create(order *entity.Order) error {
	return r.db.Create(order).Error
}

func (r *SqliteOrderRepository) FindByID(orderId string) (*entity.Order, error) {

	order := &entity.Order{}
	err := r.db.First(order, orderId).Error
	return order, err
}

func (r *SqliteOrderRepository) FindAll() ([]entity.Order, error) {

	orders := []entity.Order{}
	err := r.db.Find(&orders).Error
	return orders, err
}
