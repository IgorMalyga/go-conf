package db

import (
	"go-conf/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db_connection *gorm.DB

func initDB() *gorm.DB {
	dsn := "host=localhost user=root password=password dbname=conf port=5432 sslmode=disable"
	db_connection, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db_connection
}

func GetDB() *gorm.DB {
	if db_connection == nil {
		return initDB()
	}

	return db_connection
}

func Migrate() {
	db_connection.AutoMigrate(&models.User{})
	db_connection.AutoMigrate(&models.Workspace{})
	db_connection.AutoMigrate(&models.Article{})
	db_connection.AutoMigrate(&models.Tag{})
}
