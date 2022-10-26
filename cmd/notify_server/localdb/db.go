package localdb

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Init(db_type string) *gorm.DB {
	//dbURL := "postgres://pg:pass@localhost:5432/crud"
	if db_type == "postgres" {
		db, err := gorm.Open(postgres.Open(localconf.Config.DSN), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		err = db.AutoMigrate(&models.User{}, &models.Birthday{})
		if err != nil {
			log.Fatalln(err)
		}
		//db.Where("1=1").Delete(&models.Birthday{})
		//db.Where("1=1").Delete(&models.User{})
		return db
	} else if db_type == "sqlite" {
		db, err := gorm.Open(sqlite.Open(localconf.Config.DSN), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		err = db.AutoMigrate(&models.User{}, &models.Birthday{})
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
