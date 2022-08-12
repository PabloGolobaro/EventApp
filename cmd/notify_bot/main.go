package main

import (
	"errors"
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/api"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/config"
	_ "github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/docs"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/excel_migrator"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/httputil"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	log.Println("Opening db...")
	//db, err := gorm.Open(sqlite.Open(config.Config.DSN), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(config.Config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Birthday{})
	if err != nil {
		log.Fatal(err)
	}
	config.Config.DB = db
	err = excel_migrator.GetDataFromExcel("Birthdays.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Creating router...")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	log.Println("Creating swagger...")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Creating routes...")
	v1 := r.Group("/api/v1")
	{
		v1.Use(auth())
		v1.GET("/birthdays/:id", api.GetBirthday)
		v1.GET("/birthdays/all", api.GetAllBirthdays)
	}
	log.Println("Starting server...")
	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		if authHeader != config.Config.ApiKey {
			httputil.NewError(c, http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to this operation: api_key=%s", authHeader))
			c.Abort()
		}
		c.Next()
	}
}
