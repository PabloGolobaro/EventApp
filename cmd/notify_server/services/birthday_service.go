package services

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"log"
)

type ExcelDAO interface {
	GetFromFile() (int, error)
	GetData() []models.Birthday
	SetToFile(filename string) error
}

type Birthday_DAO interface {
	Create(birthday models.Birthday) error
	Read(id uint) (models.Birthday, error)
	ReadAll() ([]models.Birthday, error)
	Update(id uint, birthday models.Birthday) error
	Delete(id uint) error
}
type BirthdayService struct {
	birth_dao Birthday_DAO
	excel     ExcelDAO
}

func NewBirthdayService(birth_dao Birthday_DAO) *BirthdayService {
	return &BirthdayService{birth_dao: birth_dao}
}
func (s *BirthdayService) Excel_to_db(filename string) error {
	s.excel = daos.NewExcelFileDAO(filename)
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
func (s *BirthdayService) Delete(id uint) error {
	err := s.birth_dao.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *BirthdayService) Update(id uint, birth models.Birthday) error {
	err := s.birth_dao.Update(id, birth)
	if err != nil {
		return err
	}
	return nil
}
func (s *BirthdayService) Post(birth models.Birthday) error {
	err := s.birth_dao.Create(birth)
	if err != nil {
		return err
	}
	return nil
}
