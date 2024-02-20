package tg_bot

import (
	tgbot_methods "animals-game/pkg/tgbot-addons"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// private message commands and something else
const (
	createAnimalCmd = "create"
	startCmd        = "start"
)

func privateHandler(tgBot TgBot, update *tgbotapi.Update) {
	if update.Message != nil {
		privateMessageHandler(tgBot, update.Message)
	}
}

func privateMessageHandler(tgBot TgBot, msg *tgbotapi.Message) {
	if msg.IsCommand() {
		switch msg.Command() {
		case startCmd:
			StartHandler(tgBot, msg)
		case createAnimalCmd:
			CreateAnimalHandler(tgBot, msg)
		}
	}
}

func CreateAnimalHandler(tgBot TgBot, msg *tgbotapi.Message) {

}

func StartHandler(tgBot TgBot, msg *tgbotapi.Message) {

	err := tgBot.repo.CreateUserWithTgId(msg.From.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New user %s succesfully added to db", msg.From.UserName)

	answerText := fmt.Sprintf("Ну здравствуй! Теперь я запомнил, что ты со мной говорил. Хочешь найти себе нового прекрасного друга?"+
		" Тогда скорее пиши %s и будем выбирать ", createAnimalCmd)

	answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
	answerMsg.ReplyToMessageID = msg.MessageID

	tgbot_methods.SendMsg(tgBot.bot, &answerMsg)
}
