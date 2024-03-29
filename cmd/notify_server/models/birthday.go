package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username         string `gorm:"column:username;unique" json:"name,omitempty"`
	PasswordHash     string `gorm:"column:passwordhash"`
	TelegramId       string `gorm:"column:telegram_id"`
	Email            string `gorm:"uniqueIndex;"`
	VerificationCode string
	Verified         bool       `gorm:"default:false"`
	Birthdays        []Birthday `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Birthday struct {
	gorm.Model
	FullName    string    `json:"full_name,omitempty" form:"fullname"`
	PhoneNumber string    `json:"phone_number,omitempty" form:"phonenumber"`
	BirthDate   time.Time `json:"birth_date" form:"birthdate" time_format:"2006-01-02"`
	UserID      uint
}
