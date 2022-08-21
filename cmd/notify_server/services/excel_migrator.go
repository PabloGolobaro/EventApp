package services

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
)

func GetDataFromExcel(filename string) error {
	service := NewBirthdayService(daos.NewBirthdayDAO())
	err := service.Excel_to_db(filename)
	if err != nil {
		return err
	}
	return nil
}
