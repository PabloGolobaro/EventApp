package middlewares

import (
	"errors"
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/httputil/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			helpers.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		if authHeader != config.Config.ApiKey {
			helpers.NewError(c, http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to this operation: api_key=%s", authHeader))
			c.Abort()
		}
		c.Next()
	}
}
