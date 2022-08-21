package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username   string     `gorm:"column:username;unique" json:"name,omitempty"`
	Password   string     `gorm:"column:password"`
	TelegramId string     `gorm:"column:telegram_id"`
	Birthdays  []Birthday `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Birthday struct {
	gorm.Model
	FullName    string    `json:"full_name,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	BirthDate   time.Time `json:"birth_date" gorm:"default:-"`
	UserID      uint
}