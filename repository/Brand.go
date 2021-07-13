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
		return nil, errors.New("category not found")
	}
}
func DelBrandById(id int) {
	var brand models.Brand
	database.DB.Raw("DELETE  FROM `brands` WHERE id = ? ", id).Scan(&brand)
}
func UpdateBrand(brand *models.Brand) error {
	brandvalue, err := GetBrandById(brand.Id)
	if err != nil {
		return errors.New("product not found")
	} else {
		brandvalue.Name = brand.Name
		brandvalue.Image = brand.Image
		database.DB.Save(&brandvalue)
		return nil
	}
}
