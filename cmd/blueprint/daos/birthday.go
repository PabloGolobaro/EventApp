package daos

import (
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/config"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/models"
	"gorm.io/gorm"
	"time"
)

type BirthdayDAO struct {
}

func NewBirthdayDAO() *BirthdayDAO {
	return &BirthdayDAO{}
}
func (dao BirthdayDAO) Get(id uint) (*models.Birthday, error) {
	var birthday models.Birthday
	err := config.Config.DB.First(&birthday, id).Error
	return &birthday, err
}
func (dao BirthdayDAO) Create(fullname string, phonenumber string, birthday time.Time, user_id uint) error {
	birth_struct := models.Birthday{
		FullName:    fullname,
		PhoneNumber: phonenumber,
		BirthDate:   birthday,
		UserID:      user_id,
	}
	result := config.Config.DB.Create(&birth_struct)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrNotImplemented
	}
	return nil
}
