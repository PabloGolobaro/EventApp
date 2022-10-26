package handlers

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
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
	address := misc.FindIPAddress()
	user := ctx.Sender()
	userDAO := daos.NewUserDAO(config.Config.DB)
	telegram_user, err := userDAO.ReadByTelegramId(strconv.Itoa(int(ctx.Sender().ID)))
	if err == gorm.ErrRecordNotFound {
		telegram_user = models.User{
			Username:   user.Username,
			Password:   "Password",
			TelegramId: strconv.Itoa(int(ctx.Sender().ID)),
			Birthdays:  nil,
		}
		err := userDAO.Create(telegram_user)
		if err != nil {
			return err
		}
		answer = fmt.Sprintf("Здравствуй %v. Похоже ты здесь впервые!\n Твои данные для входы на сайт:\n Username: %v password: %v", user.Username, user.Username, "Password\nАдрес сайта - "+address+":8080")
	} else if err != nil {
		return err
	} else {
		answer = fmt.Sprintf("Здравствуй %v.\n Твои данные для входы на сайт:\n Username: %v password: %v", user.Username, user.Username, "Password\nАдрес сайта - "+address+":8080")
	}

	return ctx.Send(answer, reply.Menu)
}
