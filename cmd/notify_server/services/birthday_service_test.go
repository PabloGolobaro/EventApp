package services

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestBirthdayService(t *testing.T) {
	//dsn := "host=db user=postgres password=pacan334 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Birthday{})
	db.AutoMigrate(&models.User{})
	db.Where("1 = 1").Delete(&models.Birthday{})
	db.Where("1 = 1").Delete(&models.User{})
	config.Config.DB = db
	birthdayService := NewBirthdayService(daos.NewBirthdayDAO())
	err = birthdayService.Excel_to_db("Birthdays.xlsx")
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
