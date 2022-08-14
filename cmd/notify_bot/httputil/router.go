package httputil

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/api"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/controllers"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/httputil/globals"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/httputil/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func NewGinRouter() *gin.Engine {
	r := gin.Default()
	log.Println("Creating swagger...")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Creating routes...")

	r.Static("/assets", "./static/assets")
	r.LoadHTMLGlob("./static/templates/*.html")

	r.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := r.Group("/")
	PublicRoutes(public)

	private := r.Group("/")
	private.Use(middlewares.AuthRequired)
	PrivateRoutes(private)

	return r
}
func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/", controllers.IndexGetHandler())

}

func PrivateRoutes(g *gin.RouterGroup) {
	v1 := g.Group("/api/v1")
	{
		//v1.Use(middlewares.Auth())
		v1.GET("/birthdays/:id", api.GetBirthday)
		v1.PUT("/birthdays/:id", api.PutBirthday)
		v1.POST("/birthdays/add", api.PostBirthday)
		v1.DELETE("/birthdays/:id", api.DeleteBirthday)
		v1.GET("/birthdays/all", api.GetAllBirthdays)
		v1.GET("/birthdays/today", api.TodaysBirthdays)
		v1.GET("/birthdays/tomorrow", api.TommorowBirthdays)
	}
	g.GET("/dashboard", controllers.DashboardGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())

}
