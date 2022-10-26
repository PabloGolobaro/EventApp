package db

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

const sqlite_db_name string = "test.db"

func Init(db_type string) *gorm.DB {
	//dbURL := "postgres://pg:pass@localhost:5432/crud"
	if db_type == "postgres" {
		db, err := gorm.Open(postgres.Open(config.Config.DSN), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		err = db.AutoMigrate(&models.Birthday{}, &models.User{})
		if err != nil {
			log.Fatalln(err)
		}
		return db
	} else if db_type == "sqlite" {
		db, err := gorm.Open(sqlite.Open(sqlite_db_name), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		err = db.AutoMigrate(&models.Birthday{}, &models.User{})
		if err != nil {
			log.Fatalln(err)
		}
		//db.Where("1=1").Delete(&models.User{})
		//db.Where("1=1").Delete(&models.Birthday{})
		return db
	} else {
		log.Fatal("Error: wrong type of DB")
		return nil
	}
}
