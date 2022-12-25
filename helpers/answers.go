package helpers

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func SquadRoar(c tele.Context) string {
	return fmt.Sprintf("🔥%v🔥 решил созвать SQUAD\n\n@Semanovik @AlexNicker @Andrey @Vyacheslov - время для доты настало 🕐🕐🕐", c.Sender().FirstName)
}
