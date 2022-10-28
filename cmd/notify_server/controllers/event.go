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

func EventGetHandler() gin.HandlerFunc {
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
func EventPostHandler() gin.HandlerFunc {
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
				"birthday": updatedBirthday,
				"user":     user,
				"content":  content_string})
		}

	}
}
func EventAddGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		c.HTML(http.StatusOK, "event.html", gin.H{
			"user": user,
		})

	}
}

func EventAddPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		user_id := session.Get(globals.UserId)

		var NewEvent models.Birthday

		err := c.Bind(&NewEvent)
		NewEvent.UserID = user_id.(uint)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		s := services.NewBirthdayService(daos.NewBirthdayDAO(localconf.Config.DB))
		err = s.Post(NewEvent)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			content_string := fmt.Sprintf("Данные успешно добавлены!")
			c.HTML(http.StatusOK, "event.html", gin.H{
				"user":     user,
				"birthday": NewEvent,
				"content":  content_string})
		}

	}
}
func EventDeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		s := services.NewBirthdayService(daos.NewBirthdayDAO(localconf.Config.DB))
		err := s.Delete(uint(id))
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
		} else {

			c.HTML(http.StatusOK, "empty.html", gin.H{
				"notification": "Успешно удалено",
				"user":         user,
			})
		}

	}
}
