package tgbot_addons

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

const (
	maxAttempsToSend = 5
	tgSendWaitTime   = 3 * time.Second
)

func SendMsg(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig) {
	for attemp := 0; attemp < maxAttempsToSend; attemp++ {
		_, err := bot.Send(msg)
		if err == nil {
			return
		}

		fmt.Printf("Something went wrong while try to send msg: %s", err)
		time.Sleep(tgSendWaitTime)
	}
}
