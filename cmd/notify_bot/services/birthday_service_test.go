package services

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestBirthdayService(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Birthday{})
	db.Where("1 = 1").Delete(&models.Birthday{})
	config.Config.DB = db
	birthdayService := NewBirthdayService(daos.NewBirthdayDAO(), daos.NewExcelFileDAO("Birthdays.xlsx"))
	err = birthdayService.Excel_to_db()
	if err != nil {
		log.Fatal(err)
	}
	birthdays, err := birthdayService.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, birthday := range birthdays {
		fmt.Println(birthday.FullName, birthday.BirthDate.Format("02.01.2006"))
	}
}
