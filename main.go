package main

import (
	"fmt"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"os"
	"sams_helper/iternal/side_methods"
	"sams_helper/scrap"
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

	b.Handle("тест", test)

	b.Start()
}

//TODO Сделать нормально

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
	response, err := side_methods.AllCurrencies()
	if err != nil {
		return c.Send("что-то пошло не так")
	}
	return c.Send(response)
}

func test(c tele.Context) error {
	usd, err := scrap.StealUSD("https://quote.rbc.ru/ticker/59111")
	kzt, err := scrap.StealKZT("https://www.google.com/finance/quote/RUB-KZT")
	euro, err := scrap.StealEuro("https://quote.rbc.ru/ticker/59090")
	ali, err := scrap.StealAli("https://alicoup.ru/currency/")
	if err != nil {
		return c.Send(err)
	}
	return c.Send(fmt.Sprintf("%.3f\n%.3f\n%.3f\n%v\n%.3f", usd, kzt, euro, err, ali))
}
