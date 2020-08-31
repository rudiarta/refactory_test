package model

import "time"

type User struct {
	ID        uint       `json:"id" gorm:"column:id;primaryKey"`
	FullName  string     `json:"full_name" gorm:"column:full_name"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "users"
}
