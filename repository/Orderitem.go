package repository

import (
	"errors"
	"finalbackend/database"
	"finalbackend/models"
)

func GetAllOrderitem() []models.Orderitem {
	var orderitem []models.Orderitem
	database.DB.Raw("SELECT * FROM `orderitems` ").Scan(&orderitem)
	return orderitem
}
func GetOrderitemById(id int) (*models.Orderitem, error) {
	var orderitem models.Orderitem
	database.DB.Raw("SELECT * FROM `orderitems` WHERE id = ? ", id).Scan(&orderitem)
	if orderitem.Id != 0 {
		return &orderitem, nil
	} else {
		return nil, errors.New("orderitem not found")
	}
}
func DelOrderitemById(id int) {
	var orderitem models.Orderitem
	database.DB.Raw("DELETE  FROM `orderitems` WHERE id = ? ", id).Scan(&orderitem)
}
func UpdateOrderitem(orderitem *models.Orderitem) error {
	orderitemvalue, err := GetOrderitemById(orderitem.Id)
	if err != nil {
		return errors.New("orderitem not found")
	} else {
		orderitemvalue.OrderId = orderitem.OrderId
		orderitemvalue.Order = orderitem.Order
		orderitemvalue.ProductId = orderitem.ProductId
		orderitemvalue.Product = orderitem.Product
		orderitemvalue.Quantity = orderitem.Quantity

		database.DB.Save(&orderitemvalue)
		return nil
	}
}
