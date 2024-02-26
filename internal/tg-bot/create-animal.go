package tg_bot

import (
	"animals-game/configs"
	"animals-game/internal/animals"
	tgbotaddons "animals-game/pkg/tgbot-addons"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"time"
)

var stickersConfig configs.StickersConfig

func createAnimalHandler(tgBot TgBot, msg *tgbotapi.Message) {

	tgBot.chatsCache[msg.Chat.ID] = tgbotaddons.SetCache(creatingAnimalStarted, time.Now(), animals.NewAnimal()) // set cache for chat

	answerText := "Окей, давай начнем с выбора животного! Отправь мне стикер из этого стикерпака с животным, которое хочешь завести!" //todo: create more variants

	answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
	answerMsg.ReplyToMessageID = msg.MessageID

	slog.Info("", stickersConfig)

	stickerID := tgbotapi.FileID(stickersConfig.StickerExample)
	stickerMsg := tgbotapi.NewSticker(msg.Chat.ID, stickerID)

	tgbotaddons.SendMsg(tgBot.bot, answerMsg)
	tgbotaddons.SendMsg(tgBot.bot, stickerMsg)
}
