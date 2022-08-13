package api

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/models"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetBirthday godoc
// @Summary Retrieves birthday based on given ID
// @Produce json
// @Param id path integer true "Birthday ID"
// @Success 200 {object} models.Birthday
// @Router /birthdays/{id} [get]
// @Security ApiKeyAuth
func GetBirthday(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO())
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
// @Router /birthdays/all [get]
// @Security ApiKeyAuth
func GetAllBirthdays(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO())
	if birthdays, err := s.GetAll(); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, birthdays)
	}
}

// PostBirthday godoc
// @Summary Puts birthday based on given data
// @Produce json
// @Param birth_struct body models.Birthday true "Birthday structure"
// @Success 200 {object} models.Birthday
// @Router /birthdays/add [post]
// @Security ApiKeyAuth
func PostBirthday(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO())
	var birthday models.Birthday
	if err := c.BindJSON(&birthday); err != nil {
		return
	}
	if err := s.Put(birthday); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, birthday)
	}
}

// PutBirthday godoc
// @Summary Updates birthday based on given ID and data
// @Produce json
// @Param id path integer true "Birthday ID"
// @Param birth_struct body models.Birthday true "Birthday structure"
// @Success 200 {object} models.Birthday
// @Router /birthdays/{id} [put]
// @Security ApiKeyAuth
func PutBirthday(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO())
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
// @Summary Deletes birthday based on given ID
// @Produce json
// @Param id path integer true "Birthday ID"
// @Success 200
// @Router /birthdays/{id} [delete]
// @Security ApiKeyAuth
func DeleteBirthday(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := s.Delete(uint(id)); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "Success")
	}
}
