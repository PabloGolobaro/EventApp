package main

import (
	"fmt"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/apis"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/config"
	_ "github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	config.Config.DB = db
	log.Println("Creating router...")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	log.Println("Creating swagger...")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Creating routes...")
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id", apis.GetUser)
	}
	log.Println("Starting server...")
	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
	// load application configurations

}
