package reply

import tele "gopkg.in/telebot.v3"

var (
	// Universal markup builders.
	Menu = &tele.ReplyMarkup{ResizeKeyboard: true, RemoveKeyboard: true, OneTimeKeyboard: true}

	// Reply buttons.
	BtnHelp              = Menu.Text("ℹ Помощь")
	BtnExcelFile         = Menu.Text("⚙ Добавить данные через Excel файл")
	BtnPay               = Menu.Text("⚙ Купить/продлить доступ")
	BtnShowAll           = Menu.Text("🤖 Вывести все напоминания")
	BtnShowToday         = Menu.Text("🎉 Сегодня")
	BtnShowTommorow      = Menu.Text("✌ Завтра")
	BtnShowMonth         = Menu.Text("✨ Месяц")
	BtnStartNotification = Menu.Text("🤖 Включить оповещение")
	BtnStopNotification  = Menu.Text("🤖 Выключить оповещение")
)

func init() {
	Menu.Reply(
		Menu.Row(BtnHelp, BtnExcelFile),
		Menu.Row(BtnShowAll),
		Menu.Row(BtnShowToday, BtnShowTommorow, BtnShowMonth),
		Menu.Row(BtnStartNotification, BtnStopNotification),
		Menu.Row(BtnPay),
	)
}
