package models

import "time"

type Model struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key;column:id"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}
type User struct {
	Model
	FirstName string `gorm:"column:first_name" json:"first_name,omitempty"`
	LastName  string `gorm:"column:last_name" json:"last_name,omitempty"`
	Address   string `gorm:"column:addres" json:"addres,omitempty"`
	Email     string `gorm:"column:email" json:"email,omitempty"`
}
