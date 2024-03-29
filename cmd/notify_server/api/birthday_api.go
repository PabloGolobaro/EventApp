package api

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/core"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetBirthday godoc
// @Summary Retrieves birthdays based on given ID
// @Produce json
// @Param id path integer true "Birthday ID"
// @Success 200 {object} models.Birthday
// @Router /birthdays/{id} [get]
// @Security ApiKeyAuth
func GetBirthday(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO(localconf.Config.DB))
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if birth, err := s.GetById(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, birth)
	}
}

// GetAllBirthdays godoc
// @Summary Retrieves all birthdays
// @Produce json
// @Success 200 {objects} models.Birthday
// @Router /birthdays [get]
// @Security ApiKeyAuth
func GetAllBirthdays(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO(localconf.Config.DB))
	if birthdays, err := s.GetAll(); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, birthdays)
	}
}

// PostBirthday godoc
// @Summary Puts birthdays based on given data
// @Produce json
// @Param birth_struct body models.Birthday true "Birthday structure"
// @Success 200 {object} models.Birthday
// @Router /birthdays [post]
// @Security ApiKeyAuth
func PostBirthday(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO(localconf.Config.DB))
	var birthday models.Birthday
	if err := c.BindJSON(&birthday); err != nil {
		return
	}
	if err := s.Post(birthday); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, birthday)
	}
}

// PutBirthday godoc
// @Summary Updates birthdays based on given ID and data
// @Produce json
// @Param id path integer true "Birthday ID"
// @Param birth_struct body models.Birthday true "Birthday structure"
// @Success 200 {object} models.Birthday
// @Router /birthdays/{id} [put]
// @Security ApiKeyAuth
func PutBirthday(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO(localconf.Config.DB))
	var birthday models.Birthday
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := c.BindJSON(&birthday); err != nil {
		return
	}
	if err := s.Update(uint(id), birthday); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, birthday)
	}
}

// DeleteBirthday godoc
// @Summary Deletes birthdays based on given ID
// @Produce json
// @Param id path integer true "Birthday ID"
// @Success 200
// @Router /birthdays/{id} [delete]
// @Security ApiKeyAuth
func DeleteBirthday(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO(localconf.Config.DB))
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := s.Delete(uint(id)); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "Success")
	}
}
func TodaysBirthdays(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO(localconf.Config.DB))
	if birthdays, err := s.GetAll(); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		todayBirthdays := core.CheckTodayBirthdays(birthdays)
		c.JSON(http.StatusOK, todayBirthdays)
	}
}
func TommorowBirthdays(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO(localconf.Config.DB))
	if birthdays, err := s.GetAll(); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		tomorrowBirthdays := core.CheckTomorrowBirthdays(birthdays)
		c.JSON(http.StatusOK, tomorrowBirthdays)
	}
}
