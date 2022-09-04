package handlers

import (
	"bytes"
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/core"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/keyboard/inline"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/keyboard/reply"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"strconv"
)

const excel_temp_file = "excel_temp.xlsx"

var HelpButton = func(ctx tele.Context) error {
	answer := fmt.Sprintf("Описание работы бота")
	return ctx.Send(answer, reply.Menu)
}
var ExcelButton = func(ctx tele.Context) error {
	answer := fmt.Sprintf("Отправьте файл в формате Excel")
	return ctx.Send(answer)
}

var ShowAllButton = func(ctx tele.Context) error {
	userDAO := daos.NewUserDAO(config.Config.DB)
	user, err := userDAO.ReadByTelegramId(strconv.Itoa(int(ctx.Sender().ID)))
	if err != nil {
		return err
	}
	var birthdays []models.Birthday
	birthdays, err = findBirthdaysCache(birthdays, user.ID, userDAO)
	if err != nil {
		return err
	}
	page, ok := config.Config.PageMap.Map[ctx.Sender().ID]
	if !ok {
		config.Config.PageMap.Map[ctx.Sender().ID] = 0
	}
	answer := fmt.Sprintf("Ваши напоминания:\n")
	for i, birthday := range birthdays {
		if i >= 0+page*10 && i < (page+1)*10 {
			answer += fmt.Sprintf("%d. %v: %s %s\n", i+1, birthday.FullName, birthday.BirthDate.Format("02.01.2006"), birthday.PhoneNumber)
		}
		if i >= (page+1)*10 {
			break
		}

	}
	return ctx.EditOrSend(answer, inline.Selector)
}

var ParseExcelDocument = func(ctx tele.Context) error {

	file := ctx.Message().Document.File
	file_name := ctx.Message().Document.FileName
	fmt.Println(file_name)
	readCloser, err := ctx.Bot().File(&file)
	if err != nil {
		log.Println(err)
	}
	defer readCloser.Close()
	var buff bytes.Buffer
	_, err = buff.ReadFrom(readCloser)
	if err != nil {
		return err
	}
	openFile, err := os.OpenFile(excel_temp_file, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer openFile.Close()
	_, err = buff.WriteTo(openFile)
	if err != nil {
		return err
	}
	user, err := daos.NewUserDAO(config.Config.DB).ReadByTelegramId(strconv.Itoa(int(ctx.Sender().ID)))
	if err != nil {
		return err
	}
	excel := daos.NewExcelFileDAO(openFile.Name())
	_, err = excel.GetFromFile(user.ID)
	if err != nil {
		return err
	}
	birthdays := excel.GetData()
	birthdayDAO := daos.NewBirthdayDAO(config.Config.DB)
	for _, birthday := range birthdays {
		err := birthdayDAO.Create(birthday)
		if err != nil {
			return err
		}
	}
	err = os.Remove(excel_temp_file)
	if err != nil {
		log.Println(err)
	}
	answer := fmt.Sprintf("Файл сохранен в базу дынных.")
	return ctx.EditOrSend(answer)
}

var ShowTodayButton = func(ctx tele.Context) error {
	userDAO := daos.NewUserDAO(config.Config.DB)
	user, err := userDAO.ReadByTelegramId(strconv.Itoa(int(ctx.Sender().ID)))
	if err != nil {
		return err
	}
	var birthdays []models.Birthday
	birthdays, err = findBirthdaysCache(birthdays, user.ID, userDAO)
	if err != nil {
		return err
	}
	birthdays = core.CheckTodayBirthdays(birthdays)
	answer := fmt.Sprintf("Ваши напоминания:\n")
	for i, birthday := range birthdays {
		answer += fmt.Sprintf("%d. %v: %s %s\n", i+1, birthday.FullName, birthday.BirthDate.Format("02.01.2006"), birthday.PhoneNumber)
	}
	return ctx.EditOrSend(answer)
}
var ShowTommorowButton = func(ctx tele.Context) error {
	userDAO := daos.NewUserDAO(config.Config.DB)
	user, err := userDAO.ReadByTelegramId(strconv.Itoa(int(ctx.Sender().ID)))
	if err != nil {
		return err
	}
	var birthdays []models.Birthday
	birthdays, err = findBirthdaysCache(birthdays, user.ID, userDAO)
	if err != nil {
		return err
	}
	birthdays = core.CheckTomorrowBirthdays(birthdays)
	answer := fmt.Sprintf("Ваши напоминания:\n")
	for i, birthday := range birthdays {
		answer += fmt.Sprintf("%d. %v: %s %s\n", i+1, birthday.FullName, birthday.BirthDate.Format("02.01.2006"), birthday.PhoneNumber)
	}
	return ctx.EditOrSend(answer)
}
var ShowMonthButton = func(ctx tele.Context) error {
	userDAO := daos.NewUserDAO(config.Config.DB)
	user, err := userDAO.ReadByTelegramId(strconv.Itoa(int(ctx.Sender().ID)))
	if err != nil {
		return err
	}
	var birthdays []models.Birthday
	birthdays, err = findBirthdaysCache(birthdays, user.ID, userDAO)
	if err != nil {
		return err
	}
	birthdays = core.CheckMonthBirthdays(birthdays)
	answer := fmt.Sprintf("Ваши напоминания:\n")
	for i, birthday := range birthdays {
		answer += fmt.Sprintf("%d. %v: %s %s\n", i+1, birthday.FullName, birthday.BirthDate.Format("02.01.2006"), birthday.PhoneNumber)
	}
	return ctx.EditOrSend(answer)
}
