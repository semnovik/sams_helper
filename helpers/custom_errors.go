package helpers

import (
	"fmt"
	"github.com/prometheus/common/log"
	tele "gopkg.in/telebot.v3"
)

func SendErrorOnBackend(context tele.Context, someErr error) {
	err := context.Send(fmt.Sprintf("๐๐๐ะงัะพ-ัะพ ะฟะพัะปะพ ะฝะตั ัะฐะบ: %v ๐๐๐", someErr))
	if err != nil {
		log.Error(err)
	}
}
