package daos

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"gorm.io/gorm"
)

type BirthdayDAO struct {
	DB *gorm.DB
}

func NewBirthdayDAO(db *gorm.DB) *BirthdayDAO {
	return &BirthdayDAO{DB: db}
}

func (dao BirthdayDAO) Update(id uint, birthday models.Birthday) error {
	var birth models.Birthday
	db := dao.DB.First(&birth, id)
	if db.Error != nil {
		return db.Error
	}
	birth.FullName = birthday.FullName
	birth.PhoneNumber = birthday.PhoneNumber
	birth.BirthDate = birthday.BirthDate
	dao.DB.Save(&birth)
	return nil
}

func (dao BirthdayDAO) Delete(id uint) error {
	var birth models.Birthday
	db := dao.DB.First(&birth, id)
	if db.Error != nil {
		return db.Error
	}
	dao.DB.Delete(&birth)
	return nil
}

func (dao BirthdayDAO) Create(birthday models.Birthday) error {
	result := dao.DB.Create(&birthday)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrNotImplemented
	}
	return nil
}
func (dao BirthdayDAO) Read(id uint) (models.Birthday, error) {
	var birth_struct models.Birthday
	db := dao.DB.First(&birth_struct, id)
	if db.Error != nil {
		return birth_struct, db.Error
	}
	return birth_struct, nil
}
func (dao BirthdayDAO) ReadAll() ([]models.Birthday, error) {
	var birthdays []models.Birthday
	db := dao.DB.Find(&birthdays)
	fmt.Println(db.RowsAffected)
	if db.Error != nil {
		return birthdays, db.Error
	}
	return birthdays, nil
}
