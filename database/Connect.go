package database

import (
	"finalbackend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
<<<<<<< HEAD
	database, err := gorm.Open(mysql.Open("tung:12345678@/keyholic_test"), &gorm.Config{})
=======
<<<<<<< HEAD
	database, err := gorm.Open(mysql.Open("tung:12345678@/keyholic"), &gorm.Config{})
=======
	database, err := gorm.Open(mysql.Open("root:@/keyholic"), &gorm.Config{})
>>>>>>> 28e18b7fe29baedf032232cbcd9cc754ca9c0425
>>>>>>> 8f18e162030db55adf7da1a8a82b487f204e39f8
	if err != nil {
		panic("couldn't connect to the database")
	}
	DB = database
	DB.AutoMigrate(&models.Charge{}, &models.Product{}, &models.Brand{}, &models.Category{}, &models.Detailproduct{}, &models.Imageproduct{}, &models.Order{}, &models.Orderitem{}, &models.Trueproduct{}, &models.User{})
	return DB
}
