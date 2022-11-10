package controllers

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil/globals"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil/helpers"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/misc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetSignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"content": "Please enter all the fields",
		})
	}
}

// [...] SignUp User
func SignUpUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := &models.SignUpInput{}
		payload.Name = ctx.PostForm("username")
		payload.Password = ctx.PostForm("password")
		payload.PasswordConfirm = ctx.PostForm("passwordConfirm")
		payload.Email = ctx.PostForm("email")
		//if err := ctx.MultipartForm(&payload); err != nil {
		//	ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		//	return
		//}

		if payload.Password != payload.PasswordConfirm {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
			return
		}

		hashedPassword, err := helpers.HashPassword(payload.Password)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
			return
		}
		newUser := &models.User{
			Username:     payload.Name,
			Email:        strings.ToLower(payload.Email),
			PasswordHash: hashedPassword,
			Verified:     false,
		}

		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
			return
		}
		// Generate Verification Code
		code := misc.RandString(10)

		verification_code := helpers.Encode(code)
		// Update User in Database
		newUser.VerificationCode = verification_code
		localconf.Config.DB.Save(newUser)

		var firstName = newUser.Username

		if strings.Contains(firstName, " ") {
			firstName = strings.Split(firstName, " ")[1]
		}

		// ? Send Email
		emailData := helpers.EmailData{
			URL:      "https://" + localconf.Config.Domain + "/verifyemail/" + code,
			Username: firstName,
			Subject:  "Your account verification code",
		}

		helpers.SendEmail(newUser, &emailData)

		message := "We sent an email with a verification code to " + newUser.Email + ".\nPlease Check your Email."
		ctx.HTML(http.StatusOK, "send_code.html", gin.H{
			"content": message,
		})
	}

}
func VerifyEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Params.ByName("verificationCode")
		verification_code := helpers.Encode(code)

		var updatedUser models.User
		result := localconf.Config.DB.First(&updatedUser, "verification_code = ?", verification_code)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid verification code or user doesn't exists"})
			return
		}

		if updatedUser.Verified {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User already verified"})
			return
		}

		updatedUser.VerificationCode = ""
		updatedUser.Verified = true
		localconf.Config.DB.Save(&updatedUser)
		//ctx.Redirect(http.StatusFound, "https://"+localconf.Config.Domain+"/login")
		message := "Your Email is verified! Now you can Log In!"
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"success": message,
		})
	}

}
func AttachBot() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		telegramId := ctx.Query("id")
		userId := session.Get(globals.UserId)
		user := session.Get(globals.Userkey)

		var updatedUser models.User
		result := localconf.Config.DB.First(&updatedUser, userId)
		if result.Error != nil {
			ctx.HTML(http.StatusBadRequest, "send_code.html", gin.H{"content": "Invalid verification code or user doesn't exists"})
			return
		}

		if updatedUser.TelegramId != "" {
			ctx.HTML(http.StatusBadRequest, "send_code.html", gin.H{"content": "Telegram Bot is already connected to your account!"})
			return
		}

		updatedUser.TelegramId = telegramId
		localconf.Config.DB.Save(&updatedUser)
		message := "Telegram Bot is successfully connected to your account!"
		ctx.HTML(http.StatusOK, "send_code.html", gin.H{
			"content": message,
			"user":    user,
		})
	}

}
