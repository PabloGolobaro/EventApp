package services

import "github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"

type User_DAO interface {
	Create(user models.User) error
	Read(id uint) (models.User, error)
	ReadAll() ([]models.User, error)
	Update(id uint, user models.User) error
	Delete(id uint) error
	GetBirthdays(user *models.User) ([]models.Birthday, error)
	GetPaginatedBirthdays(user *models.User, pagination *models.Pagination) ([]models.Birthday, error)
}
type UserService struct {
	user_dao User_DAO
}

func NewUserService(user_dao User_DAO) *UserService {
	return &UserService{user_dao: user_dao}
}

func (s *UserService) GetById(id uint) (models.User, error) {
	one, err := s.user_dao.Read(id)
	if err != nil {
		return one, err
	}
	return one, nil
}
func (s *UserService) GetAll() ([]models.User, error) {
	users, err := s.user_dao.ReadAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *UserService) Post(birth models.User) error {
	err := s.user_dao.Create(birth)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserService) GetAllUserBirthdays(user *models.User) ([]models.Birthday, error) {
	birthdays, err := s.user_dao.GetBirthdays(user)
	if err != nil {
		return nil, err
	}
	return birthdays, nil
}
func (s *UserService) GetPaginatedBirthdays(user *models.User, pagination *models.Pagination) ([]models.Birthday, error) {
	birthdays, err := s.user_dao.GetPaginatedBirthdays(user, pagination)
	if err != nil {
		return nil, err
	}
	return birthdays, nil
}
