package daos

import (
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/config"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/models"
)

type UserDAO struct {
}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao UserDAO) Get(id uint) (*models.User, error) {

	var user models.User
	err := config.Config.DB.Where("id = ?", id).
		First(&user).
		Error

	return &user, err
}
