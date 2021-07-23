package repository

import (
	"finalbackend/database"
	"finalbackend/models"
)

func SavePayment(charge *models.Charge) (err error) {
	if err = database.DB.Create(&charge).Error; err != nil {
		return err
	}
	return nil
}
