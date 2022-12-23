package main

import (
	"gopkg.in/telebot.v3/middleware"
	"log"
	"os"
	"sams_helper/iternal/side_methods"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	b.Use(middleware.Logger())

	b.Handle(tele.OnText, handleCustom)
	b.Handle("праздник", holidays)
	b.Handle("курс", currencyAll)

	b.Start()
}

//TODO Сделать слой транспорта

func handleCustom(c tele.Context) error {

	var (
		_    = c.Sender()
		text = c.Text()
	)
	if strings.Contains(text, "привет") {
		return c.Send("принял")
	}
	return c.Send("не принял")
}

func holidays(c tele.Context) error {
	response := side_methods.GetListOfHolidays()
	return c.Send(response)
}

func currencyAll(c tele.Context) error {
	response := side_methods.GetCurrentCurrencyUSD() + "\n" + side_methods.GetAliCurrency()
	return c.Send(response)
}
