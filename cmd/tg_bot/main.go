package main

import (
	"flag"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/db"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/handlers"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/keyboard/inline"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/keyboard/reply"
	"github.com/PabloGolobaro/go-notify-project/cmd/tg_bot/misc"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"time"
)

func main() {
	db_type := flag.String("db", "postgres", "Type of DB to use.")
	flag.Parse()

	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	pref := tele.Settings{
		Token:  config.Config.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Opening db...")
	config.Config.DB = db.Init(*db_type)

	b.Use(middleware.AutoRespond())

	b.Handle(tele.OnText, func(c tele.Context) error {
		markup := b.NewMarkup()
		markup.Inline(
			markup.Row(markup.URL("Visit", "https://google.com"), markup.URL("Visit too", "https://github.com")),
		)

		return c.EditOrSend("Menu", markup)
	})

	b.Handle(tele.OnDocument, handlers.ParseExcelDocument)
	b.Handle("/start", handlers.StartCommand)
	b.Handle(&reply.BtnHelp, handlers.HelpButton)
	b.Handle(&reply.BtnExcelFile, handlers.ExcelButton)
	b.Handle(&reply.BtnShowAll, handlers.ShowAllButton)
	b.Handle(&reply.BtnShowToday, handlers.ShowTodayButton)
	b.Handle(&reply.BtnShowTommorow, handlers.ShowTommorowButton)
	b.Handle(&reply.BtnShowMonth, handlers.ShowMonthButton)
	b.Handle(&inline.BtnNext, handlers.NextButton)
	b.Handle(&inline.BtnPrev, handlers.PrevButton)
	b.Handle(&inline.BtnCancel, handlers.CancelButton)
	b.Handle(&reply.BtnPay, handlers.PayButton)
	b.Handle(tele.OnCheckout, handlers.PreCheckout)
	b.Handle(tele.OnPayment, handlers.GetPayment)
	b.Handle(&reply.BtnStartNotification, handlers.StartNotificationFunc)
	b.Handle(&reply.BtnStopNotification, handlers.StopNotificationFunc)

	//	markup := b.NewMarkup()
	//	markup.Inline(
	//		markup.Row(markup.URL("Visit", "https://google.com"), markup.URL("Visit too", "https://github.com")),
	//	)
	b.Handle(tele.OnQuery, handlers.InlinePhotos)
	log.Println("Starting bot...")
	err = misc.SendToAdmins(b)
	if err != nil {
		log.Println(err)
	}
	b.Start()
}
