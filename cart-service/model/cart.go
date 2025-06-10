package model

type CartItem struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	UserID    uint `json:"useer_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
