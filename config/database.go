package config

import (
	"github.com/emejotaw/graphql-api/internal/database/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {

	dsn := "database.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = autoMigrate(db)

	if err != nil {
		panic(err)
	}

	return db
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Order{}, &entity.Product{})
}
