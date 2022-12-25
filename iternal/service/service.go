package service

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"sams_helper/helpers"
	"sams_helper/helpers/buttons"
	"sams_helper/iternal/side_methods"
	"sams_helper/scrap"
	"strings"
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
	selector := buttons.NewInlineSelector()

	switch c.Message().Chat.ID {
	case -1001268955342:
		selector.Markup.Inline(
			selector.Markup.Row(selector.HolidaysBtn(), selector.CurrencyBtn()),
			selector.Markup.Row(selector.SquadCallBtn()),
		)
	default:
		selector.Markup.Inline(
			selector.Markup.Row(selector.HolidaysBtn(), selector.CurrencyBtn()),
		)
	}

	return c.Send(fmt.Sprintf("%v, что тебе от меня надо?", c.Sender().FirstName), selector.Markup)
}

func (s Service) HandleCallback(c tele.Context) error {
	str := c.Callback().Data
	switch {
	case strings.Contains(str, "holiday"):
		c.Send(fmt.Sprintf("%v хочет посмотреть праздники", c.Sender().FirstName))

		holidays, err := side_methods.GetListOfHolidays()
		if err != nil {
			helpers.SendErrorOnBackend(c, err)
		}
		return c.Send(holidays)
	case strings.Contains(str, "currency"):
		c.Send(fmt.Sprintf("%v хочет узнать курс валют", c.Sender().FirstName))

		currency, err := side_methods.AllCurrencies()
		if err != nil {
			helpers.SendErrorOnBackend(c, err)
		}
		return c.Send(currency)

	case strings.Contains(str, "squad call"):
		return c.Send(helpers.SquadRoar(c))

		//case strings.Contains(str, "data"):
		//	return c.Send("test")
	}
	return nil
}
