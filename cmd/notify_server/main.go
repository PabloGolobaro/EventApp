package main

import (
	"flag"
	"fmt"
	//_ "github.com/GoAdminGroup/go-admin/adapter/gin"               // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
	//_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // Import the sql driver
	//_ "github.com/GoAdminGroup/themes/adminlte"                    // Import the theme
	"github.com/gin-gonic/autotls"

	_ "github.com/PabloGolobaro/go-notify-project/cmd/notify_server/docs"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localdb"
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
	if err := localconf.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	db_type := flag.String("db", "postgres", "Type of DB to use.")
	flag.Parse()
	log.Println("Opening db...")
	localconf.Config.DB = localdb.Init(*db_type)
	log.Println("Creating router...")
	r := httputil.NewGinRouter()
	log.Println("Starting server...")
	log.Fatal(autotls.Run(r, localconf.Config.Domain))
}

//eng := engine.Default()

// GoAdmin global configuration, can also be imported as a json file.
//cfg := config.Config{
//	Databases: config.DatabaseList{
//		"default": {
//			MaxIdleCon: 50,
//			MaxOpenCon: 150,
//			File:       "./admin.db",
//			Driver:     db.DriverSqlite,
//		},
//	},
//	UrlPrefix: "admin", // The url prefix of the website.
//	// Store must be set and guaranteed to have write access, otherwise new administrator users cannot be added.
//	Store: config.Store{
//		Path:   "./uploads",
//		Prefix: "uploads",
//	},
//	Language: language.EN,
//}

// Add configuration and plugins, use the Use method to mount to the web framework.
//_ = eng.AddConfig(&cfg).AddGenerators(admin_panel.Generators).Use(r)
