package api

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO(config.Config.DB))
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.GetById(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetAllUserBirthdays godoc
// @Summary Retrieves all birthdays from given user
// @Produce json
// @Success 200 {objects} models.Birthday
// @Router /users/{id}/all [get]
// @Security ApiKeyAuth
func GetAllUserBirthdays(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO(config.Config.DB))
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if birthdays, err := s.GetAllUserBirthdays(uint(id)); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, birthdays)
	}
}

// PostUser godoc
// @Summary Post user based on given data
// @Produce json
// @Param user_struct body models.User true "User structure"
// @Success 200 {object} models.User
// @Router /users [post]
// @Security ApiKeyAuth
func PostUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO(config.Config.DB))
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	if err := s.Post(user); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
