package repository

import (
	"errors"
	"finalbackend/database"
	"finalbackend/models"
)

func GetAllDetailproduct() []models.Detailproduct {
	var detailproduct []models.Detailproduct
	database.DB.Raw("SELECT * FROM `detailproducts` ").Scan(&detailproduct)
	return detailproduct
}
func GetDetailproductById(id int) (*models.Detailproduct, error) {
	var detailproduct models.Detailproduct
	database.DB.Raw("SELECT * FROM `detailproducts` WHERE id = ? ", id).Scan(&detailproduct)
	if detailproduct.Id != 0 {
		return &detailproduct, nil
	} else {
		return nil, errors.New("detailproduct not found")
	}
}
func DelDetailproductById(id int) {
	var detailproduct models.Detailproduct
	database.DB.Raw("DELETE  FROM `detailproducts` WHERE id = ? ", id).Scan(&detailproduct)
}
func UpdateDetailproduct(detailproduct *models.Detailproduct) error {
	detailproductvalue, err := GetCategoryById(detailproduct.Id)
	if err != nil {
		return errors.New("detailproduct not found")
	} else {
		detailproductvalue.Name = detailproduct.Name
		database.DB.Save(&detailproductvalue)
		return nil
	}
}
