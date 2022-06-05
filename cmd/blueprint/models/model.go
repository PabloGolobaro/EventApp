package models

import (
	"time"
)

type Model struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key;column:id"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}
