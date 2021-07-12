package repository

import (
	"errors"
	"finalbackend/database"
	"finalbackend/models"
)

func GetAllImageproduct() []models.Imageproduct {
	var imageproduct []models.Imageproduct
	database.DB.Raw("SELECT * FROM `imageproducts` ").Scan(&imageproduct)
	return imageproduct
}
func GetImageproductById(id int) (*models.Imageproduct, error) {
	var imageproduct models.Imageproduct
	database.DB.Raw("SELECT * FROM `imageproducts` WHERE id = ? ", id).Scan(&imageproduct)
	if imageproduct.Id != 0 {
		return &imageproduct, nil
	} else {
		return nil, errors.New("imageproduct not found")
	}
}
func DelImageproductById(id int) {
	var imageproduct models.Imageproduct
	database.DB.Raw("DELETE  FROM `imageproducts` WHERE id = ? ", id).Scan(&imageproduct)
}
func UpdateImageproduct(imageproduct *models.Imageproduct) error {
	imageproductvalue, err := GetImageproductById(imageproduct.Id)
	if err != nil {
		return errors.New("imageproduct not found")
	} else {
		imageproductvalue.Product = imageproduct.Product
		imageproductvalue.ProductId = imageproduct.ProductId
		imageproductvalue.Image = imageproduct.Image

		database.DB.Save(&imageproductvalue)
		return nil
	}
}
