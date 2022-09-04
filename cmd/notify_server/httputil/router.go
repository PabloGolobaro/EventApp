package httputil

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/api"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/controllers"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil/globals"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil/helpers"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"html/template"
	"log"
)

func NewGinRouter() *gin.Engine {
	r := gin.Default()
	log.Println("Creating swagger...")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Creating routes...")
	r.SetFuncMap(template.FuncMap{
		"short": helpers.Shorten_date,
	})
	r.Static("/assets", "./static/assets")
	r.LoadHTMLGlob("./static/templates/*.html")

	r.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	private := r.Group("/")
	private.Use(middlewares.AuthRequired)
	PrivateRoutes(private)

	public := r.Group("/")
	PublicRoutes(public)

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
		birthdays := v1.Group("/birthdays")
		{
			birthdays.PUT("/:id", api.PutBirthday)
			birthdays.POST("/", api.PostBirthday)
			birthdays.GET("/:id", api.GetBirthday)
			birthdays.DELETE("/:id", api.DeleteBirthday)
			birthdays.GET("/", api.GetAllBirthdays)
			birthdays.GET("/today", api.TodaysBirthdays)
			birthdays.GET("/tomorrow", api.TommorowBirthdays)
		}
		users := v1.Group("/users")
		{
			users.GET("/:id", api.GetUser)
			users.GET("/:id/all", api.GetAllUserBirthdays)
			users.POST("/", api.PostUser)
		}
	}
	g.GET("/dashboard", controllers.DashboardGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())
}
