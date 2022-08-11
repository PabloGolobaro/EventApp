package excel_migrator

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/services"
)

func GetDataFromExcel(filename string) error {
	service := services.NewBirthdayService(daos.NewBirthdayDAO(), daos.NewExcelFileDAO(filename))
	err := service.Excel_to_db()
	if err != nil {
		return err
	}
	return nil
}
