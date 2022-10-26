package handlers

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/httputil/helpers"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/keyboard/inline"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

var CancelButton = func(ctx tele.Context) error {
	delete(config.Config.PageMap.Map, ctx.Sender().ID)
	answer := fmt.Sprintf("Возврат в меню!")
	return ctx.Edit(answer)
}
var NextButton = func(ctx tele.Context) error {
	userDAO := daos.NewUserDAO(config.Config.DB)
	user, err := userDAO.ReadByTelegramId(strconv.Itoa(int(ctx.Sender().ID)))
	if err != nil {
		return err
	}

	var birthdays []models.Birthday
	birthdays, err = findBirthdaysCache(birthdays, user, userDAO)
	if err != nil {
		return err
	}
	page, ok := config.Config.PageMap.Map[ctx.Sender().ID]
	if !ok {
		config.Config.PageMap.Map[ctx.Sender().ID] = 0
	} else {
		if page != len(birthdays)/10 {
			config.Config.PageMap.Map[ctx.Sender().ID] += 1
			page += 1
		} else {
			return nil
		}
	}
	answer := fmt.Sprintf("Ваши напоминания:\n")
	for i, birthday := range birthdays {
		if i >= 0+page*10 && i < (page+1)*10 {
			answer += fmt.Sprintf("%d. %v: %s %s\n", i+1, birthday.FullName, birthday.BirthDate.Format("02.01.2006"), birthday.PhoneNumber)
		}

	}
	return ctx.Edit(answer, inline.Selector)
}
var PrevButton = func(ctx tele.Context) error {
	userDAO := daos.NewUserDAO(config.Config.DB)
	user, err := userDAO.ReadByTelegramId(strconv.Itoa(int(ctx.Sender().ID)))
	if err != nil {
		return err
	}

	var birthdays []models.Birthday
	birthdays, err = findBirthdaysCache(birthdays, user, userDAO)
	if err != nil {
		return err
	}
	page, ok := config.Config.PageMap.Map[ctx.Sender().ID]
	if !ok {
		config.Config.PageMap.Map[ctx.Sender().ID] = 0
	} else {
		if page != 0 {
			config.Config.PageMap.Map[ctx.Sender().ID] -= 1
			page -= 1
		} else {
			return nil
		}

	}
	answer := fmt.Sprintf("Ваши напоминания:\n")
	for i, birthday := range birthdays {
		if i >= 0+page*10 && i < (page+1)*10 {
			answer += fmt.Sprintf("%d. %v: %s %s\n", i+1, birthday.FullName, birthday.BirthDate.Format("02.01.2006"), birthday.PhoneNumber)
		}

	}
	return ctx.Edit(answer, inline.Selector)
}

func findBirthdaysCache(birthdays []models.Birthday, user models.User, userDAO *daos.UserGorm) ([]models.Birthday, error) {
	var err error
	cache, ok := config.Config.Cache.M[user.ID]
	if !ok {
		config.Config.Cache.M[user.ID] = make([]models.Birthday, 0)
		birthdays, err = userDAO.GetBirthdays(&user)
		if err != nil {
			return birthdays, err
		}
		birthdays = helpers.Sort_birthdays(birthdays)
		config.Config.Cache.M[user.ID] = append(config.Config.Cache.M[user.ID], birthdays...)
	} else {
		birthdays = cache
	}
	return birthdays, nil
}
