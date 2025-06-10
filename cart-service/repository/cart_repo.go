package repository

import (
	"github.com/prashant-bhilwar/order-processing-system/cart-service/database"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/model"
)

func AddToCart(item *model.CartItem) error {
	return database.DB.Create(item).Error
}

func RemoveFromCart(userID uint, productID uint) error {
	return database.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&model.CartItem{}).Error
}

func GetCartItemsByUserID(userID uint) ([]model.CartItem, error) {
	var items []model.CartItem
	err := database.DB.Where("user_id = ?", userID).Find(&items).Error
	return items, err
}
