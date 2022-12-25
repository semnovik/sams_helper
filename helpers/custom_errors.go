package helpers

import (
	"fmt"
	"github.com/prometheus/common/log"
	tele "gopkg.in/telebot.v3"
)

func SendErrorOnBackend(context tele.Context, someErr error) {
	err := context.Send(fmt.Sprintf("💀💀💀Что-то пошло нет так: %v 💀💀💀", someErr))
	if err != nil {
		log.Error(err)
	}
}
