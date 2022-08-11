package api

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/daos"
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
	s := services.NewBirthdayService(daos.NewBirthdayDAO(), daos.NewExcelFileDAO("Birthdays.xlsx"))
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
// @No Params
// @Success 200 {objects} models.Birthday
// @Router /birthdays/all [get]
// @Security ApiKeyAuth
func GetAllBirthdays(c *gin.Context) {
	s := services.NewBirthdayService(daos.NewBirthdayDAO(), daos.NewExcelFileDAO("Birthdays.xlsx"))
	if birthdays, err := s.GetAll(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, birthdays)
	}
}
