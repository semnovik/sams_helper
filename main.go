package main

import (
	"log"
	"os"
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

	b.Handle(tele.OnText, handleCustom)

	b.Start()
}

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
