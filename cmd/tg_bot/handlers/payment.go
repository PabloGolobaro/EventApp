package handlers

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

var PayButton = func(ctx tele.Context) error {
	invoice := tele.Invoice{
		Title:       "Подписка",
		Description: "Купить подписку на месяц",
		Payload:     "Token",
		Currency:    "RUB",
		Prices: []tele.Price{
			tele.Price{
				Label:  "Недельная",
				Amount: 15000,
			}, tele.Price{
				Label:  "Месячная",
				Amount: 45000,
			},
		},
		Token:               "381764678:TEST:43374",
		Data:                "",
		Photo:               nil,
		PhotoSize:           0,
		Start:               "",
		Total:               60000,
		MaxTipAmount:        1500,
		SuggestedTipAmounts: []int{500, 1000, 1500},
		NeedName:            true,
		NeedPhoneNumber:     true,
		NeedEmail:           true,
		NeedShippingAddress: false,
		SendPhoneNumber:     false,
		SendEmail:           false,
		Flexible:            false,
	}
	options := tele.SendOptions{}
	send, err := invoice.Send(ctx.Bot(), ctx.Sender(), &options)
	if err != nil {
		return err
	}
	fmt.Println(send)
	return nil
}
var PreCheckout = func(ctx tele.Context) error {
	err := ctx.Bot().Accept(ctx.PreCheckoutQuery())
	if err != nil {
		return err
	}
	return nil
}

var GetPayment = func(ctx tele.Context) error {
	payment := ctx.Message().Payment
	fmt.Println(payment)
	return nil
}
