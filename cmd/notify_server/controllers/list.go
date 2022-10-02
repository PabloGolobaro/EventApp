package controllers

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/core"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil/globals"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil/helpers"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		readByUsername, err := daos.NewUserDAO(localconf.Config.DB).ReadByUsername(user.(string))
		if err != nil {
			return
		}
		birthdays := readByUsername.Birthdays
		birthdays = helpers.Sort_birthdays(birthdays)
		pagination := helpers.GeneratePaginationFromRequest(c, len(birthdays))
		offset := (pagination.Page - 1) * pagination.Limit
		if offset+pagination.Limit > len(birthdays) {
			birthdays = birthdays[offset:]
		} else {
			birthdays = birthdays[offset : offset+pagination.Limit]
		}

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"content":    "Ваши события",
			"user":       user,
			"birthdays":  birthdays,
			"pagination": pagination,
			"path":       c.FullPath(),
		})
	}
}
func MonthGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		readByUsername, err := daos.NewUserDAO(localconf.Config.DB).ReadByUsername(user.(string))
		if err != nil {
			return
		}
		birthdays := readByUsername.Birthdays
		birthdays = core.CheckMonthBirthdays(birthdays)
		birthdays = helpers.Sort_birthdays(birthdays)
		pagination := helpers.GeneratePaginationFromRequest(c, len(birthdays))
		offset := (pagination.Page - 1) * pagination.Limit
		if offset+pagination.Limit > len(birthdays) {
			birthdays = birthdays[offset:]
		} else {
			birthdays = birthdays[offset : offset+pagination.Limit]
		}
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"content":    "Ваши дни рождения в этом месяце",
			"user":       user,
			"birthdays":  birthdays,
			"pagination": pagination,
			"path":       c.FullPath(),
		})
	}
}
func TodayGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		readByUsername, err := daos.NewUserDAO(localconf.Config.DB).ReadByUsername(user.(string))
		if err != nil {
			return
		}
		birthdays := readByUsername.Birthdays
		birthdays = core.CheckTodayBirthdays(birthdays)
		birthdays = helpers.Sort_birthdays(birthdays)

		pagination := helpers.GeneratePaginationFromRequest(c, len(birthdays))
		offset := (pagination.Page - 1) * pagination.Limit
		if offset+pagination.Limit > len(birthdays) {
			birthdays = birthdays[offset:]
		} else {
			birthdays = birthdays[offset : offset+pagination.Limit]
		}
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"content":    "Ваши дни рождения сегодня",
			"user":       user,
			"birthdays":  birthdays,
			"pagination": pagination,
			"path":       c.FullPath(),
		})
	}
}
func TomorrowGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		readByUsername, err := daos.NewUserDAO(localconf.Config.DB).ReadByUsername(user.(string))
		if err != nil {
			return
		}
		birthdays := readByUsername.Birthdays
		birthdays = core.CheckTomorrowBirthdays(birthdays)
		birthdays = helpers.Sort_birthdays(birthdays)
		pagination := helpers.GeneratePaginationFromRequest(c, len(birthdays))
		offset := (pagination.Page - 1) * pagination.Limit
		if offset+pagination.Limit > len(birthdays) {
			birthdays = birthdays[offset:]
		} else {
			birthdays = birthdays[offset : offset+pagination.Limit]
		}
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"content":    "Ваши дни рождения на завтра",
			"user":       user,
			"birthdays":  birthdays,
			"pagination": pagination,
			"path":       c.FullPath(),
		})
	}
}
