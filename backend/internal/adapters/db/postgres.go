package db

import (
	"fmt"
	"log"

	"github.com/jakkapan-a/fiber-next-ecommerce/config"
	"github.com/jakkapan-a/fiber-next-ecommerce/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSL,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// migrateDatabase(db)

	log.Println("Database connection established successfully")
	return db
}

func migrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.Role{},
		&domain.Permission{},
		&domain.User{},
		&domain.Category{},
		&domain.Product{},
		&domain.ProductImage{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Transaction{},
	)

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")

}
