package model

import "time"

type Product struct {
	ID         int        `json:"id" gorm:"column:id;primaryKey"`
	Name       string     `json:"name" gorm:"column:name"`
	Variant    string     `json:"variant" gorm:"column:variant"`
	MerchantID int        `json:"merchant_id" gorm:"column:merchant_id"`
	Price      int        `json:"price" gorm:"column:price"`
	Status     bool       `json:"status" gorm:"column:status"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Product) TableName() string {
	return "products"
}
