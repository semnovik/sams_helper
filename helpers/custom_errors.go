package helpers

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func SendErrorOnBackend(context tele.Context, err error) error {
	return context.Send(fmt.Sprintf("💀💀💀Что-то пошло нет так: %v 💀💀💀", err))
}
