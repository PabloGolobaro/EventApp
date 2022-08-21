package daos

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"gorm.io/gorm"
)

type UserGorm struct {
}

func (u UserGorm) Create(user models.User) error {
	result := config.Config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrNotImplemented
	}
	return nil
}

func (u UserGorm) Read(id uint) (models.User, error) {
	var user_struct models.User
	db := config.Config.DB.First(&user_struct, id)
	if db.Error != nil {
		return user_struct, db.Error
	}
	return user_struct, nil
}
func (u UserGorm) ReadByUsername(username string) (models.User, error) {
	var user_struct models.User
	db := config.Config.DB.Where("username = ?", username).First(&user_struct)
	if db.Error != nil {
		return user_struct, db.Error
	}
	return user_struct, nil
}

func (u UserGorm) ReadAll() ([]models.User, error) {
	var users []models.User
	db := config.Config.DB.Find(&users)
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
func (u UserGorm) GetBirthdays(id uint) ([]models.Birthday, error) {
	var birthdays []models.Birthday
	user, err := u.Read(id)
	if err != nil {
		return nil, err
	}
	err = config.Config.DB.Order("birth_date").Model(&user).Association("Birthdays").Find(&birthdays)
	if err != nil {
		return nil, err
	}
	return birthdays, nil

}
func NewUserDAO() *UserGorm {
	return &UserGorm{}
}
