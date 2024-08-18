package database

import (
	// "go-fiber-gorm/models/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connection := "rull:rull@tcp(127.0.0.1:3306)/go-restapi?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed Connect to DB")
	}

	// db.AutoMigrate(&entity.User{})

	DB = db
	log.Println("DB Conncect")
}