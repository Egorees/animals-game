package tg_bot

import (
	tgbotmethods "animals-game/pkg/tgbot-addons"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"time"
)

const (
	techSupportTg           = "@mregores"
	userTgAlreadyExistError = "pq: duplicate key value violates unique constraint \"users_telegram_id_key\""
)

// private message commands and something else
const (
	createAnimal = "create"
	start        = "start"
	cancel       = "cancel"
)

func privateHandler(tgBot TgBot, update *tgbotapi.Update) {
	if update.Message != nil {
		privateMessageHandler(tgBot, update.Message)
	}
}

func privateMessageHandler(tgBot TgBot, msg *tgbotapi.Message) {
	if msg.IsCommand() {
		switch msg.Command() {
		case start:

			startHandler(tgBot, msg)
		case createAnimal:
			createAnimalHandler(tgBot, msg)
		case cancel: //todo: realise cancel handler
		}
	}

	switch tgBot.chatsCache[msg.Chat.ID].Command {
	case mainMenu:

	default:
		slog.Info("Unknown command:", tgBot.chatsCache[msg.Chat.ID].Command)
	}

	if msg.Sticker != nil {
		//testing todo: remove this
		answerText := fmt.Sprintf("Твое животное: %s!", stickersConfig.GetAnimaTypeByStickerID(msg.Sticker.FileUniqueID))
		answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
		answerMsg.ReplyToMessageID = msg.MessageID
		tgbotmethods.SendMsg(tgBot.bot, answerMsg)
	}
}

func startHandler(tgBot TgBot, msg *tgbotapi.Message) {

	tgBot.chatsCache[msg.Chat.ID] = tgbotmethods.SetCache(mainMenu, time.Now(), nil) // set cache for chat

	err := tgBot.repo.CreateUserWithTgId(msg.From.ID)

	var answerText string

	if err != nil {
		slog.Error("Oh no!", err)

		if err.Error() == userTgAlreadyExistError {
			answerText = "Ну какой старт! Уже знакомились же... "
		} else {
			answerText = fmt.Sprintf("Какая-то непонятная ошибка, напиши пожалуйста %s", techSupportTg)
		}
	} else {
		slog.Info("New user successfully added to db", msg.From.UserName)

		answerText = fmt.Sprintf("Ну здравствуй! Теперь я запомнил, что ты со мной говорил. Хочешь найти себе нового прекрасного друга?"+
			" Тогда скорее пиши /%s и будем выбирать ", createAnimal)
	}

	answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
	answerMsg.ReplyToMessageID = msg.MessageID

	tgbotmethods.SendMsg(tgBot.bot, answerMsg)
}
