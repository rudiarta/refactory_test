package model

type Cart struct {
	ID        int `json:"id" gorm:"column:id;primaryKey`
	UserID    int `json:"user_id" gorm:"column:user_id"`
	ProductID int `json:"product_id" gorm:"column:product_id"`
}

func (Cart) TableName() string {
	return "carts"
}
