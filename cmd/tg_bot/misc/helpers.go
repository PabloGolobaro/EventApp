package misc

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/config"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
)

func SendToAdmins(bot *tele.Bot) error {
	for _, admin := range config.Config.Admins {
		id, err := strconv.Atoi(admin)
		if err != nil {
			log.Println(fmt.Errorf("Ошибка отправки оповещения админам: %v", err))
			continue
		}
		_, err = bot.Send(tele.ChatID(id), "Бот запущен")
		if err != nil {
			return fmt.Errorf("Ошибка отправки оповещения админам: %v", err)
		}
	}
	return nil
}
