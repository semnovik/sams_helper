package main

import (
	"log"
	"os"
	"sams_helper/iternal/handler"
	"sams_helper/iternal/service"
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

	srv := service.NewService("sema_help_me_pls")
	handler.InitHandlers(b, srv)

	b.Start()
}

// TODO: единая точка входа для url
// TODO: конфигурация переменных окружения
