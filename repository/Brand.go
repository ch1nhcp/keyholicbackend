package repository

import (
	"errors"
	"finalbackend/database"
	"finalbackend/models"
)

func GetAllBrand() []models.Brand {
	var brand []models.Brand
	database.DB.Raw("SELECT * FROM `brands` ").Scan(&brand)
	return brand
}
func GetBrandById(id int) (*models.Brand, error) {
	var brand models.Brand
	database.DB.Raw("SELECT * FROM `brands` WHERE id = ? ", id).Scan(&brand)
	if brand.Id != 0 {
		return &brand, nil
	} else {
		return nil, errors.New("Category not found")
	}
}
func DelBrandById(id int) {
	var brand models.Brand
	database.DB.Raw("DELETE  FROM `brands` WHERE id = ? ", id).Scan(&brand)
}
func UpdateBrand(brand *models.Brand) error {
	brandd, err := GetBrandById(brand.Id)
	if err != nil {
		return errors.New("Product not found")
	} else {
		brandd.Name = brand.Name
		brandd.Image = brand.Image
		database.DB.Save(&brandd)
		return nil
	}
}
