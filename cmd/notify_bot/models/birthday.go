package models

import (
	"time"
)

type User struct {
	Model
	FirstName  string     `gorm:"column:first_name" json:"first_name,omitempty"`
	LastName   string     `gorm:"column:last_name" json:"last_name,omitempty"`
	TelegramId string     `gorm:"column:telegram_id"`
	Birthday   []Birthday `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Birthday struct {
	Model
	FullName    string    `json:"full_name,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	BirthDate   time.Time `json:"birth_date"`
	//UserID      uint
}
