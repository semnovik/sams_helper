package service

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"sams_helper/helpers"
	"sams_helper/iternal/side_methods"
	"sams_helper/scrap"
)

type Service struct {
	Some string
}

func NewService(some string) *Service {
	return &Service{Some: some}
}

func (s Service) Holidays(c tele.Context) error {
	response, err := side_methods.GetListOfHolidays()
	if err != nil {
		helpers.SendErrorOnBackend(c, err)
	}
	return c.Send(response)
}

func (s Service) CurrencyAll(c tele.Context) error {
	response, err := side_methods.AllCurrencies()
	if err != nil {
		helpers.SendErrorOnBackend(c, err)
	}
	return c.Send(response)
}

func (s Service) Test(c tele.Context) error {
	usd, err := scrap.StealUSD("https://quote.rbc.ru/ticker/59111")
	kzt, err := scrap.StealKZT("https://www.google.com/finance/quote/RUB-KZT")
	euro, err := scrap.StealEuro("https://quote.rbc.ru/ticker/59090")
	ali, err := scrap.StealAli("https://alicoup.ru/currency/")
	if err != nil {
		return c.Send(err)
	}
	return c.Send(fmt.Sprintf("%.3f\n%.3f\n%.3f\n%v\n%.3f", usd, kzt, euro, err, ali))
}

func (s Service) InlineCall(c tele.Context) error {
	return c.Send("че надо", &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{{{
			Text: "Праздник сегодня",
			Data: "holiday",
		},
			{
				Text: "Курсы валют",
				Data: "currency",
			},
		}},
	})
}

func (s Service) HandleCallback(c tele.Context) error {
	switch c.Callback().Data {
	case "holiday":
		holidays, err := side_methods.GetListOfHolidays()
		if err != nil {
			helpers.SendErrorOnBackend(c, err)
		}
		return c.Send(holidays)
	case "currency":
		currency, err := side_methods.AllCurrencies()
		if err != nil {
			helpers.SendErrorOnBackend(c, err)
		}
		return c.Send(currency)
	}
	return nil
}
