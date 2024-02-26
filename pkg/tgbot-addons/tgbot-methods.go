package tgbot_addons

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"time"
)

const (
	maxAttemptsToSend = 5
	tgSendWaitTime    = 3 * time.Second
)

func SendMsg(bot *tgbotapi.BotAPI, msg tgbotapi.Chattable) {
	err := error(nil)
	for attempt := 0; attempt < maxAttemptsToSend; attempt++ {
		_, err = bot.Send(msg)
		if err == nil {
			return
		}

		time.Sleep(tgSendWaitTime)
	}
	slog.Error("Something went wrong while try to send msg:", "error", err)
}
