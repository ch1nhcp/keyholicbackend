package repository

import (
	"errors"
	"finalbackend/database"
	"finalbackend/models"
)

func GetAllOrder() []models.Order {
	var order []models.Order
	database.DB.Raw("SELECT * FROM `orders` ").Scan(&order)
	return order
}
func GetOrderById(id int) (*models.Order, error) {
	var order models.Order
	database.DB.Raw("SELECT * FROM `orders` WHERE id = ? ", id).Scan(&order)
	if order.Id != 0 {
		return &order, nil
	} else {
		return nil, errors.New("order not found")
	}
}
func DelOrderById(id int) {
	var order models.Order
	database.DB.Raw("DELETE  FROM `orders` WHERE id = ? ", id).Scan(&order)
}
func UpdateOrder(order *models.Order) error {
	ordervalue, err := GetOrderById(order.Id)
	if err != nil {
		return errors.New("order not found")
	} else {
		ordervalue.UserId = order.UserId
		ordervalue.User = order.User
		ordervalue.Name = order.Name
		ordervalue.Phone = order.Phone
		ordervalue.Address = order.Address
		ordervalue.TotalProducts = order.TotalProducts
		ordervalue.Price = order.Price
		ordervalue.Sale = order.Sale
		ordervalue.Ship = order.Ship
		ordervalue.TotalPrice = order.TotalPrice

		database.DB.Save(&ordervalue)
		return nil
	}
}
