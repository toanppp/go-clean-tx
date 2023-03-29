package database

import (
	"github.com/toanppp/go-clean-tx/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&domain.Wallet{}); err != nil {
		panic("failed to migrate the schema")
	}

	return db
}
