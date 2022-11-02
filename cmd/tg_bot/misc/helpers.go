package misc

import (
	"encoding/json"
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/config"
	tele "gopkg.in/telebot.v3"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func SendToAdmins(bot *tele.Bot) error {
	for _, admin := range config.Config.Admins {
		id, err := strconv.Atoi(admin)
		if err != nil {
			log.Println(fmt.Errorf("Ошибка отправки оповещения админам: %v", err))
			continue
		}
		_, err = bot.Send(tele.ChatID(id), "Бот запущен по адресу https://"+config.Config.Domain)
		if err != nil {
			return fmt.Errorf("Ошибка отправки оповещения админам: %v", err)
		}
	}
	return nil
}

type IP struct {
	Query string
}

func FindIPAddress() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}
