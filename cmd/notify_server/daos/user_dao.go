package daos

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"gorm.io/gorm"
)

type UserGorm struct {
	DB *gorm.DB
}

func (u UserGorm) Create(user models.User) error {
	result := u.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrNotImplemented
	}
	return nil
}

func (u UserGorm) Read(id uint) (models.User, error) {
	var user_struct models.User
	db := u.DB.First(&user_struct, id)
	if db.Error != nil {
		return user_struct, db.Error
	}
	return user_struct, nil
}
func (u UserGorm) ReadByUsername(username string) (models.User, error) {
	var user_struct models.User
	db := u.DB.Where("username = ?", username).Preload("Birthdays").First(&user_struct)
	if db.Error != nil {
		return user_struct, db.Error
	}
	return user_struct, nil
}
func (u UserGorm) ReadByTelegramId(telegram_id string) (models.User, error) {
	var user_struct models.User
	db := u.DB.Where("telegram_id = ?", telegram_id).First(&user_struct)
	if db.Error != nil {
		return user_struct, db.Error
	}
	return user_struct, nil
}

func (u UserGorm) ReadAll() ([]models.User, error) {
	var users []models.User
	db := u.DB.Find(&users)
	fmt.Println(db.RowsAffected)
	if db.Error != nil {
		return users, db.Error
	}
	return users, nil
}

func (u UserGorm) Update(id uint, birthday models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserGorm) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}
func (u UserGorm) GetBirthdays(user *models.User) ([]models.Birthday, error) {
	var birthdays []models.Birthday
	err := u.DB.Order("birth_date").Model(user).Association("Birthdays").Find(&birthdays)
	if err != nil {
		return nil, err
	}
	return birthdays, nil
}
func (u UserGorm) GetPaginatedBirthdays(user *models.User, pagination *models.Pagination) ([]models.Birthday, error) {
	var birthdays []models.Birthday
	offset := (pagination.Page - 1) * pagination.Limit
	err := u.DB.Limit(pagination.Limit).Offset(offset).Order("birth_date").Model(user).Association("Birthdays").Find(&birthdays)
	if err != nil {
		return nil, err
	}
	return birthdays, nil
}
func NewUserDAO(db *gorm.DB) *UserGorm {
	return &UserGorm{DB: db}
}
