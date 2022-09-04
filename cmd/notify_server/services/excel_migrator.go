package services

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
)

func GetDataFromExcel(telegram_id uint, filename string) error {
	service := NewBirthdayService(daos.NewBirthdayDAO(config.Config.DB))
	err := service.Excel_to_db(telegram_id, filename)
	if err != nil {
		return err
	}
	return nil
}
