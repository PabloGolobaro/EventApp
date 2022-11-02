package handlers

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil/helpers"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/keyboard/reply"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/misc"
	tele "gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"strconv"
)

var StartCommand = func(ctx tele.Context) error {
	var answer string
	//address := misc.FindIPAddress()
	user := ctx.Sender()
	userDAO := daos.NewUserDAO(config.Config.DB)
	telegram_user, err := userDAO.ReadByTelegramId(strconv.Itoa(int(ctx.Sender().ID)))
	if err == gorm.ErrRecordNotFound {
		randPass := misc.RandString(8)
		hashPassword, err := helpers.HashPassword(randPass)
		if err != nil {
			return err
		}
		telegram_user = models.User{
			Username:     user.Username,
			PasswordHash: hashPassword,
			TelegramId:   strconv.Itoa(int(ctx.Sender().ID)),
			Birthdays:    nil,
		}
		err = userDAO.Create(telegram_user)
		if err != nil {
			return err
		}
		answer = fmt.Sprintf("Здравствуй %v. Похоже ты здесь впервые!\nТвои данные для входы на сайт:\nUsername: %v password: %v\nАдрес сайта - https://%v", user.Username, user.Username, randPass, config.Config.Domain)
	} else if err != nil {
		return err
	} else {
		answer = fmt.Sprintf("Здравствуй %v.\nТвои данные для входы на сайт:\nUsername: %v\nАдрес сайта - https://%v", user.Username, user.Username, config.Config.Domain)
	}

	return ctx.Send(answer, reply.Menu)
}
