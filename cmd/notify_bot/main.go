package main

import (
	"flag"
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/db"
	_ "github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/docs"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/excel_migrator"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/httputil"
	"log"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email yeu344@gmail.com

// @license.name MIT
// @license.url https://github.com/PAbloGolobaro/go-notify-project/LICENSE

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	db_type := flag.String("db", "postgres", "Type of DB to use.")
	flag.Parse()
	log.Println("Opening db...")
	config.Config.DB = db.Init(*db_type)
	log.Println("Get Excel data...")
	err := excel_migrator.GetDataFromExcel("Birthdays.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Creating router...")
	r := httputil.NewGinRouter()
	log.Println("Starting server...")
	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}
