package httputil

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func NewGinRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	log.Println("Creating swagger...")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Creating routes...")
	v1 := r.Group("/api/v1")
	{
		v1.Use(Auth())
		v1.GET("/birthdays/:id", api.GetBirthday)
		v1.PUT("/birthdays/:id", api.PutBirthday)
		v1.POST("/birthdays/add", api.PostBirthday)
		v1.DELETE("/birthdays/:id", api.DeleteBirthday)
		v1.GET("/birthdays/all", api.GetAllBirthdays)
	}
	return r
}
