package database

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"workshop/models"
)

var DB *gorm.DB

func ConnectAndMigrate() error {
	dsn := strings.TrimSpace(os.Getenv("DB_DSN"))
	if dsn == "" {
		// host-run: 127.0.0.1  |  docker-compose: mysql
		dsn = "workshop:workshoppass@tcp(127.0.0.1:3306)/workshop?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("connect db: %w", err)
	}

	// (optional) ตั้งค่า connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get sql db: %w", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	DB = db

	if err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Promotion{},
		&models.Cart{},
		&models.CartItem{},
	); err != nil {
		return fmt.Errorf("auto migrate: %w", err)
	}
	return nil
}
