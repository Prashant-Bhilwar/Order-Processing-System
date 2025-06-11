package repository

import (
	"github.com/prashant-bhilwar/order-processing-system/order-service/database"
	"github.com/prashant-bhilwar/order-processing-system/order-service/model"
)

func CreateOrder(order *model.Order) error {
	return database.DB.Create(order).Error
}

func GetOrdersByUserID(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := database.DB.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
