package handler

import (
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"sams_helper/iternal/service"
)

func InitHandlers(b *tele.Bot, srv *service.Service) {
	b.Use(middleware.Logger())

	b.Handle("праздник", srv.Holidays)
	b.Handle("курс", srv.CurrencyAll)
	b.Handle("тест", srv.Test)
	b.Handle("бот", srv.InlineCall)
	b.Handle(tele.OnCallback, srv.HandleCallback)
}
