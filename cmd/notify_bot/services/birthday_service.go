package services

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/models"
	"log"
)

type ExcelDAO interface {
	GetFromFile() (int, error)
	GetData() []models.Birthday
	SetToFile(filename string) error
}

type birthday_DAO interface {
	Create(birthday models.Birthday) error
	Read(id uint) (models.Birthday, error)
	ReadAll() ([]models.Birthday, error)
}
type BirthdayService struct {
	birth_dao birthday_DAO
	excel     ExcelDAO
}

func NewBirthdayService(birth_dao birthday_DAO, excel ExcelDAO) *BirthdayService {
	return &BirthdayService{birth_dao: birth_dao, excel: excel}
}
func (s *BirthdayService) Excel_to_db() error {
	_, err := s.excel.GetFromFile()
	if err != nil {
		return err
	}
	for i, birthday := range s.excel.GetData() {
		err := s.birth_dao.Create(birthday)
		if err != nil {
			log.Printf("Error at %d element creating to DB!", i)
			return err
		}
	}
	return nil
}
func (s *BirthdayService) GetAll() ([]models.Birthday, error) {
	all, err := s.birth_dao.ReadAll()
	if err != nil {
		return nil, err
	}
	return all, nil
}
func (s *BirthdayService) GetById(id uint) (models.Birthday, error) {
	one, err := s.birth_dao.Read(id)
	if err != nil {
		return one, err
	}
	return one, nil
}
