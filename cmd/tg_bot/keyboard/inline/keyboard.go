package inline

import tele "gopkg.in/telebot.v3"

var (
	Selector = &tele.ReplyMarkup{}
	// Inline buttons.
	//
	// Pressing it will cause the client to
	// send the bot a callback.
	//
	// Make sure Unique stays unique as per button kind
	// since it's required for callback routing to work.
	//
	BtnPrev   = Selector.Data("⬅", "prev")
	BtnNext   = Selector.Data("➡", "next")
	BtnCancel = Selector.Data("Отмена", "cancel")
)

func init() {
	Selector.Inline(
		Selector.Row(BtnPrev, BtnNext),
		Selector.Row(BtnCancel),
	)
}
