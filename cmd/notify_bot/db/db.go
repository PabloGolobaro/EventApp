package db

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Init(db_type string) *gorm.DB {
	//dbURL := "postgres://pg:pass@localhost:5432/crud"
	if db_type == "postgres" {
		db, err := gorm.Open(postgres.Open(config.Config.DSN), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		err = db.AutoMigrate(&models.Birthday{})
		if err != nil {
			log.Fatalln(err)
		}
		db.Where("1=1").Delete(&models.Birthday{})
		return db
	} else if db_type == "sqlite" {
		db, err := gorm.Open(sqlite.Open(config.Config.DSN), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		err = db.AutoMigrate(&models.Birthday{})
		if err != nil {
			log.Fatalln(err)
		}
		db.Delete(&models.Birthday{}).Where("1=1")
		return db
	} else {
		log.Fatal("Error: wrong type of DB")
		return nil
	}
}
