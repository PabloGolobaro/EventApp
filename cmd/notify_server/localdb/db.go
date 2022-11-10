package localdb

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

const sqlite_name string = "test.db"

func Init(db_type string) *gorm.DB {
	//dbURL := "postgres://pg:pass@localhost:5432/crud"
	if db_type == "postgres" {
		dsn := fmt.Sprintf(
			"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
			localconf.Config.DBHost,
			localconf.Config.DBUserName,
			localconf.Config.DBUserPassword,
			localconf.Config.DBName,
			localconf.Config.DBPort,
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
		db, err := gorm.Open(sqlite.Open(sqlite_name), &gorm.Config{})
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
