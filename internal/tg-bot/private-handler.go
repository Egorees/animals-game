package tg_bot

import (
	tgbotaddons "animals-game/pkg/tgbot-addons"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
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
		return
	}

	switch tgBot.chatsCache[msg.Chat.ID].Command {
	case mainMenu:
		// todo: think how answering on usual msg
	case waitAnimalType:
		setAnimalType(tgBot, msg)
	case waitAnimalName:
		setAnimalName(tgBot, msg)
	case waitAcceptAnimalInfo:
		acceptAnimalFeatures(tgBot, msg)
	default:
		slog.Info("Unknown command type:", tgBot.chatsCache[msg.Chat.ID].Command)
	}
}

func startHandler(tgBot TgBot, msg *tgbotapi.Message) {

	tgBot.chatsCache[msg.Chat.ID] = tgbotaddons.SetCache(mainMenu, nil) // set cache for chat

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

	tgbotaddons.SendMsg(tgBot.bot, answerMsg)
}
