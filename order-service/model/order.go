package model

import "time"

type Order struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"` // pending, confirmed, cancelled
	CreatedAt  time.Time `json:"created_at"`
	ProductID  uint      `json:"product_id"`
}
