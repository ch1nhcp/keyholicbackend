package database

import (
	"finalbackend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	database, err := gorm.Open(mysql.Open("tung:12345678@/keyholic_test"), &gorm.Config{})

	if err != nil {
		panic("couldn't connect to the database")
	}
	DB = database
	DB.AutoMigrate(&models.Charge{}, &models.Product{}, &models.Brand{}, &models.Category{}, &models.Detailproduct{}, &models.Imageproduct{}, &models.Order{}, &models.Orderitem{}, &models.Trueproduct{}, &models.User{})
	return DB
}
