package main

import (
	"flag"
	"fmt"
	api "github.com/PabloGolobaro/go-notify-project/cmd/notify_server/api/smtp_api"
	_ "github.com/PabloGolobaro/go-notify-project/cmd/notify_server/docs"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localdb"
	"github.com/gin-gonic/autotls"
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
	if err := localconf.LoadConfig("."); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	db_type := flag.String("db", "postgres", "Type of DB to use.")
	env_type := flag.String("env", "server", "Type of environment.")
	flag.Parse()
	log.Println("Opening db...")
	localconf.Config.DB = localdb.Init(*db_type)
	if *env_type == "local" {
		localconf.Config.Domain = "localhost"
		localconf.Config.DBHost = "127.0.0.1"
	}
	newConfiguration := api.NewConfiguration()
	apiClient := api.NewAPIClient(newConfiguration)
	localconf.Config.API = apiClient
	log.Println("Creating router...")
	r := httputil.NewGinRouter()
	log.Println("Starting server...")
	go httputil.RedirectToHTTPS(localconf.Config.Domain, "443")

	if *env_type == "local" {
		log.Fatal(r.RunTLS(":443", "./cert/cert.pem", "./cert/key.pem"))
	} else {
		log.Fatal(autotls.Run(r, localconf.Config.Domain))
	}

}
