package services

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil/helpers"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestUserService(t *testing.T) {
	dsn := "host=localhost user=postgres password=pacan334 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Birthday{}, &models.User{})
	if err != nil {
		fmt.Println(err)
		return
	}
	localconf.Config.DB = db
	//birthdayservise := NewBirthdayService(daos.NewBirthdayDAO())
	UserService := NewUserService(daos.NewUserDAO(localconf.Config.DB))
	//err = UserService.Post(models.User{
	//	Username:   "Golobar",
	//	Password:   "pacan334",
	//	TelegramId: "123123123",
	//})
	//if err != nil {
	//	return
	//}

	//err = birthdayservise.Excel_to_db("Birthdays.xlsx")
	//if err != nil {
	//	return
	//}
	var users []models.User
	err = db.Model(&models.User{}).Preload("Birthdays").Find(&users).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		//err := config.Config.DB.Model(&user).Association("Birthdays").Append(&models.Birthday{
		//	FullName:    "Golobaro",
		//	PhoneNumber: "76666666666",
		//	BirthDate:   time.Now(),
		//})
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		birthdays, err := UserService.GetAllUserBirthdays(&user)
		if err != nil {
			return
		}
		birthdays = helpers.Sort_birthdays(birthdays)
		for _, birthday := range birthdays {
			fmt.Println(birthday)
		}

	}

}
