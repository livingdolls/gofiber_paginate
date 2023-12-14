package database

import (
	"gofiber-paginate/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "yurina:hirate@tcp(127.0.0.1:3306)/gopaging?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection successfull")
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.Book{})
	DB = db;
}