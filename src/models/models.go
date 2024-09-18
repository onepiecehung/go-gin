package models

import (
	"api/go-gin/src/utils/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// Model represents the base model with common fields
type Model struct {
	ID         uint           `gorm:"primary_key;autoIncrement" json:"id"`
	CreatedOn  time.Time      `gorm:"autoCreateTime" json:"createdOn"`
	ModifiedOn time.Time      `gorm:"autoUpdateTime" json:"modifiedOn"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Setup initializes the database connection
func Setup() {
	var err error

	// Updated DSN for MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DatabaseSetting.User,
		config.DatabaseSetting.Password,
		config.DatabaseSetting.Host,
		config.DatabaseSetting.Name)

	// Open connection using MySQL driver
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: config.DatabaseSetting.TablePrefix},
	})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	// Set connection pool settings
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Migrate the schema
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("Error closing database: ", err)
		return
	}
	sqlDB.Close()
}

// AutoMigrate auto migrates the database
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}
}
