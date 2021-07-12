package repository

import (
	"errors"
	"finalbackend/database"
	"finalbackend/models"
)

func GetAllCategory() []models.Category {
	var category []models.Category
	database.DB.Raw("SELECT * FROM `categories` ").Scan(&category)
	return category
}
func GetCategoryById(id int) (*models.Category, error) {
	var category models.Category
	database.DB.Raw("SELECT * FROM `categories` WHERE id = ? ", id).Scan(&category)
	if category.Id != 0 {
		return &category, nil
	} else {
		return nil, errors.New("Category not found")
	}
}
func DelCategoryById(id int) {
	var category models.Category
	database.DB.Raw("DELETE  FROM `categories` WHERE id = ? ", id).Scan(&category)
}
func UpdateCategory(category *models.Category) error {
	categoryvalue, err := GetCategoryById(category.Id)
	if err != nil {
		return errors.New("Product category not found")
	} else {
		categoryvalue.Name = category.Name
		database.DB.Save(&categoryvalue)
		return nil
	}
}
