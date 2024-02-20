package tg_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (tgBot TgBot) PrivateMessageHandler(msg *tgbotapi.Message) {
	//testing example
	answer := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
	answer.ReplyToMessageID = msg.MessageID

	for _, err := tgBot.bot.Send(answer); err != nil; {
		_, err = tgBot.bot.Send(answer)
	}
}
