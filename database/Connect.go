package database

import (
	"finalbackend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
<<<<<<< HEAD
	database, err := gorm.Open(mysql.Open("root:@/keyholic"), &gorm.Config{})
=======
	database, err := gorm.Open(mysql.Open("tung:12345678@/keyholic"), &gorm.Config{})
>>>>>>> 7b24b587b2914949ccc3377c1929859183a867ab
	if err != nil {
		panic("couldn't connect to the database")
	}
	DB = database
	DB.AutoMigrate(&models.Product{}, &models.Brand{}, &models.Category{}, &models.Detailproduct{}, &models.Imageproduct{}, &models.Order{}, &models.Orderitem{}, &models.Trueproduct{}, &models.User{})
	return DB
}
