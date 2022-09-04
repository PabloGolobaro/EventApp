package reply

import tele "gopkg.in/telebot.v3"

var (
	// Universal markup builders.
	Menu = &tele.ReplyMarkup{ResizeKeyboard: true, RemoveKeyboard: true, OneTimeKeyboard: true}

	// Reply buttons.
	BtnHelp         = Menu.Text("ℹ Помощь")
	BtnExcelFile    = Menu.Text("⚙ Добавить данные через Excel файл")
	BtnShowAll      = Menu.Text("🤖 Вывести все напоминания")
	BtnShowToday    = Menu.Text("🎉 Напоминания на сегодня")
	BtnShowTommorow = Menu.Text("✌ Напоминания на завтра")
	BtnShowMonth    = Menu.Text("✨ Напоминания на месяц")
)

func init() {
	Menu.Reply(
		Menu.Row(BtnHelp),
		Menu.Row(BtnExcelFile),
		Menu.Row(BtnShowAll),
		Menu.Row(BtnShowToday, BtnShowTommorow, BtnShowMonth),
	)
}
