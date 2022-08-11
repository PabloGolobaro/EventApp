package daos

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/models"
	"gorm.io/gorm"
)

type BirthdayDAO struct {
}

func NewBirthdayDAO() *BirthdayDAO {
	return &BirthdayDAO{}
}
func (dao BirthdayDAO) Create(birthday models.Birthday) error {
	result := config.Config.DB.Create(&birthday)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrNotImplemented
	}
	return nil
}
func (dao BirthdayDAO) Read(id uint) (models.Birthday, error) {
	var birth_struct models.Birthday
	db := config.Config.DB.First(&birth_struct, id)
	if db.Error != nil {
		return birth_struct, db.Error
	}
	return birth_struct, nil
}
func (dao BirthdayDAO) ReadAll() ([]models.Birthday, error) {
	var birthdays []models.Birthday
	db := config.Config.DB.Find(&birthdays)
	fmt.Println(db.RowsAffected)
	if db.Error != nil {
		return birthdays, db.Error
	}
	return birthdays, nil
}
