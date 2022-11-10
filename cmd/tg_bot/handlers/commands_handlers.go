package handlers

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/keyboard/reply"
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
		answer = fmt.Sprintf("Здравствуй %v. Похоже ты здесь впервые и еще не привязал свой аккаунт к боту!\nСсылка для привязки бота к твоему аккаунту на сайте: <a href=\"https://%v/attach?id=%v\">Привязать</a>", user.Username, config.Config.Domain, user.ID)
	} else if err != nil {
		return err
	} else {
		answer = fmt.Sprintf("Здравствуй %v.\nТвои данные для входы на сайт:\nUsername: %v\nАдрес сайта - https://%v", user.Username, telegram_user.Username, config.Config.Domain)
	}

	return ctx.Send(answer, reply.Menu, tele.ModeHTML)
}
