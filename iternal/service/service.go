package service

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"sams_helper/helpers"
	"sams_helper/helpers/buttons"
	"sams_helper/helpers/chats"
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

	selector.Markup.Inline(
		selector.Markup.Row(selector.HolidaysBtn(), selector.CurrencyBtn()),
	)

	switch c.Message().Chat.ID {
	case chats.SquadChat:
		selector.Markup.Inline(
			selector.Markup.Row(selector.HolidaysBtn(), selector.CurrencyBtn()),
			selector.Markup.Row(selector.SquadCallBtn()),
		)
	case chats.AdminChat, chats.AliyaChat:
		selector.Markup.Inline(
			selector.Markup.Row(selector.HolidaysBtn(), selector.CurrencyBtn()),
			selector.Markup.Row(selector.LoveBtn()),
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
		photo := &tele.Photo{
			File: tele.File{FileURL: "https://i.pinimg.com/originals/8a/8b/50/8a8b50da2bc4afa933718061fe291520.jpg"},
		}
		c.Send(photo)
		return c.Send(helpers.SquadRoarMsg(c))

	case strings.Contains(str, "honey"):
		if c.Chat().ID == chats.AliyaChat {
			c.Chat().ID = chats.AdminChat
			c.Send(helpers.HoneyMsg(c))
			c.Chat().ID = chats.AliyaChat
			c.Send(helpers.HoneyDone())
		}
		if c.Chat().ID == chats.AdminChat {
			c.Chat().ID = chats.AliyaChat
			c.Send(helpers.HoneyMsg(c))
			c.Chat().ID = chats.AdminChat
			c.Send(helpers.HoneyDone())
		}
	}
	return nil
}

func (s Service) AdminCall(c tele.Context) error {
	if c.Chat().ID != chats.AdminChat {
		return nil
	}

	if strings.Contains(c.Text(), "/sq") {
		c.Chat().ID = chats.SquadChat
		c.Send(strings.Replace(c.Text(), "/sq", "", 1))
		return nil
	}

	if strings.Contains(c.Text(), "/al") {
		c.Chat().ID = chats.AliyaChat
		c.Send(strings.Replace(c.Text(), "/al", "", 1))
		return nil
	}

	return nil
}
