package handlers

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/core"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/config"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"time"
)

var tasks map[int64]*chan bool

func init() {
	tasks = make(map[int64]*chan bool)
}

var StartNotificationFunc = func(ctx tele.Context) error {
	_, ok := tasks[ctx.Sender().ID]
	if ok {
		ctx.Send("Оповещение уже включено!")
		return nil
	}
	go func(bot *tele.Bot, ID int64) {
		ticker := time.NewTicker(12 * time.Hour)
		defer ticker.Stop()
		done := make(chan bool)
		tasks[ID] = &done
		for {
			select {
			case <-done:
				bot.Send(tele.ChatID(ID), "Оповещение завершено")
				return
			case <-ticker.C:
				err := findToday(bot, ID)
				if err != nil {
					bot.Send(tele.ChatID(ID), "Ошибка оповещения: "+err.Error())
					return
				}
				err = findTommorow(bot, ID)
				if err != nil {
					bot.Send(tele.ChatID(ID), "Ошибка оповещения: "+err.Error())
					return
				}
			}
		}
	}(ctx.Bot(), ctx.Sender().ID)
	return nil
}

var StopNotificationFunc = func(ctx tele.Context) error {

	c, ok := tasks[ctx.Sender().ID]
	if !ok {
		ctx.Send("Оповещение не было включено!")
		return nil
	} else {
		*c <- true
		close(*c)
		delete(tasks, ctx.Sender().ID)
		return nil
	}
}

func findToday(bot *tele.Bot, ID int64) error {
	userDAO := daos.NewUserDAO(config.Config.DB)
	user, err := userDAO.ReadByTelegramId(strconv.Itoa(int(ID)))
	if err != nil {
		return err
	}
	var birthdays []models.Birthday
	birthdays, err = findBirthdaysCache(birthdays, user, userDAO)
	if err != nil {
		return err
	}
	birthdays = core.CheckTodayBirthdays(birthdays)
	if len(birthdays) == 0 {
		_, err = bot.Send(tele.ChatID(ID), "Напоминаний на сегодня нет")
		return err
	}
	answer := fmt.Sprintf("Ваши напоминания на сегодня:\n")
	for i, birthday := range birthdays {
		answer += fmt.Sprintf("%d. %v: %s %s\n", i+1, birthday.FullName, birthday.BirthDate.Format("02.01.2006"), birthday.PhoneNumber)
	}
	_, err = bot.Send(tele.ChatID(ID), answer)
	return err
}
func findTommorow(bot *tele.Bot, ID int64) error {
	userDAO := daos.NewUserDAO(config.Config.DB)
	user, err := userDAO.ReadByTelegramId(strconv.Itoa(int(ID)))
	if err != nil {
		return err
	}
	var birthdays []models.Birthday
	birthdays, err = findBirthdaysCache(birthdays, user, userDAO)
	if err != nil {
		return err
	}
	birthdays = core.CheckTomorrowBirthdays(birthdays)
	if len(birthdays) == 0 {
		_, err = bot.Send(tele.ChatID(ID), "Напоминаний на завтра нет")
		return err
	}
	answer := fmt.Sprintf("Ваши напоминания на завтра:\n")
	for i, birthday := range birthdays {
		answer += fmt.Sprintf("%d. %v: %s %s\n", i+1, birthday.FullName, birthday.BirthDate.Format("02.01.2006"), birthday.PhoneNumber)
	}
	_, err = bot.Send(tele.ChatID(ID), answer)
	return err
}
