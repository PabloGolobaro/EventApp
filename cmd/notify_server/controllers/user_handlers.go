package controllers

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil/globals"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func UserGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		s := services.NewBirthdayService(daos.NewBirthdayDAO(localconf.Config.DB))
		if birthday, err := s.GetById(uint(id)); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			log.Println(err)
		} else {

			c.HTML(http.StatusOK, "event.html", gin.H{
				"birthday": birthday,
				"user":     user,
			})
		}

	}
}
func UserPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		var updatedBirthday models.Birthday

		err := c.Bind(&updatedBirthday)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		s := services.NewBirthdayService(daos.NewBirthdayDAO(localconf.Config.DB))
		err = s.Update(uint(id), updatedBirthday)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			content_string := fmt.Sprintf("Данные успешно обновлены!")
			c.HTML(http.StatusOK, "event.html", gin.H{
				"user":    user,
				"content": content_string})
		}

	}
}
