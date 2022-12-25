package helpers

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func SquadRoarMsg(c tele.Context) string {
	return fmt.Sprintf("%v решил созвать SQUAD\n\n@Semanovik @AlexNicker @Andrey @Vyacheslov\n🔥🔥🔥Время для доты🔥🔥🔥", c.Sender().FirstName)
}

func HoneyMsg(c tele.Context) string {
	var name string
	switch c.Sender().FirstName {
	case "Sema":
		name = "Сёма"
	case "Alia":
		name = "Алия"
	}
	return fmt.Sprintf("%v тебя любит💕", name)
}

func HoneyDone() string {
	return "Милость доставлена🌝"
}
