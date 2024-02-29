package tg_bot

import (
	"animals-game/internal/animals"
	tgbotaddons "animals-game/pkg/tgbot-addons"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

func createAnimalHandler(tgBot TgBot, msg *tgbotapi.Message) {

	tgBot.chatsCache[msg.Chat.ID] = tgbotaddons.SetCache(waitAnimalType, animals.NewAnimal()) // set cache for chat

	answerText := "Окей, давай начнем с выбора животного! Отправь мне стикер из этого стикерпака с животным, которое хочешь завести!" //todo: create more variants

	answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
	answerMsg.ReplyToMessageID = msg.MessageID

	stickerID := tgbotapi.FileID(stickersConfig.StickerExample)
	stickerMsg := tgbotapi.NewSticker(msg.Chat.ID, stickerID)

	tgbotaddons.SendMsg(tgBot.bot, answerMsg)
	tgbotaddons.SendMsg(tgBot.bot, stickerMsg)
}

func setAnimalType(tgBot TgBot, msg *tgbotapi.Message) {

	if msg.Sticker == nil {
		answerText := "Я вообще-то жду стикер из этого стикерпака!"
		answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
		answerMsg.ReplyToMessageID = msg.MessageID
		tgbotaddons.SendMsg(tgBot.bot, answerMsg)

		stickerID := tgbotapi.FileID(stickersConfig.StickerExample)
		stickerMsg := tgbotapi.NewSticker(msg.Chat.ID, stickerID)
		tgbotaddons.SendMsg(tgBot.bot, stickerMsg)
		return
	}

	animalType := stickersConfig.GetAnimaTypeByStickerID(msg.Sticker.FileUniqueID)
	if animalType == "" {
		answerText := "Я вообще-то жду стикер из этого стикерпака!"
		answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
		answerMsg.ReplyToMessageID = msg.MessageID
		tgbotaddons.SendMsg(tgBot.bot, answerMsg)

		stickerID := tgbotapi.FileID(stickersConfig.StickerExample)
		stickerMsg := tgbotapi.NewSticker(msg.Chat.ID, stickerID)
		tgbotaddons.SendMsg(tgBot.bot, stickerMsg)
		return
	}

	currentAnimal, ok := tgBot.chatsCache[msg.Chat.ID].OtherInfo.(*animals.Animal)

	if !ok {
		slog.Error("Something went wrong while try to set type of animal, other cache was:", currentAnimal)
		SendErrorResponse(tgBot, msg.Chat.ID)
		return
	}

	currentAnimal.Type = animalTypeFromStrToId[animalType]
	tgBot.chatsCache[msg.Chat.ID] = tgbotaddons.SetCache(waitAnimalName, currentAnimal)

	answerText := "Отлично! Теперь придумай имя своему новому другу!"
	answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
	answerMsg.ReplyToMessageID = msg.MessageID
	tgbotaddons.SendMsg(tgBot.bot, answerMsg)

}

func setAnimalName(tgBot TgBot, msg *tgbotapi.Message) {
	if msg.Text == "" {
		answerText := "Я вообще-то жду имя( Давай давай пиши!"
		answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
		answerMsg.ReplyToMessageID = msg.MessageID
		tgbotaddons.SendMsg(tgBot.bot, answerMsg)
		return
	}

	currentAnimal, ok := tgBot.chatsCache[msg.Chat.ID].OtherInfo.(*animals.Animal)

	if !ok {
		slog.Error("Something went wrong while try to set type of animal, other cache was:", currentAnimal)
		SendErrorResponse(tgBot, msg.Chat.ID)
		return
	}

	currentAnimal.Name = msg.Text
	tgBot.chatsCache[msg.Chat.ID] = tgbotaddons.SetCache(waitAcceptAnimalInfo, currentAnimal)

	answerText := fmt.Sprintf("Отлично! Ты хочешь следующее животное:\nТип: %s \nИмя: %s \nВерно? \nОтветь 'Да' или 'Нет'", animalTypeFromIdToStr[currentAnimal.Type], currentAnimal.Name)
	answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
	answerMsg.ReplyToMessageID = msg.MessageID
	tgbotaddons.SendMsg(tgBot.bot, answerMsg)
}

func acceptAnimalFeatures(tgBot TgBot, msg *tgbotapi.Message) {

	switch msg.Text {
	case "Да":
		currentAnimal, ok := tgBot.chatsCache[msg.Chat.ID].OtherInfo.(*animals.Animal)
		if !ok {
			slog.Error("Something went wrong while try to accept animal info", currentAnimal)
			SendErrorResponse(tgBot, msg.Chat.ID)
			return
		}
		var err error
		currentAnimal.OwnerId, err = tgBot.repo.GetUserIdByTgId(msg.From.ID)

		if err != nil {
			slog.Error("error while try to get userIdByTgId:", err)
			SendErrorResponse(tgBot, msg.Chat.ID)
			return
		}

		err = tgBot.repo.CreateAnimal(currentAnimal)

		if err != nil {
			slog.Error("error while try to create animal:", err)
			SendErrorResponse(tgBot, msg.Chat.ID)
			return
		}

		tgBot.chatsCache[msg.Chat.ID] = tgbotaddons.SetCache(mainMenu, nil)

		answerText := "Ура! Обращайся с ним хорошо!"
		answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
		answerMsg.ReplyToMessageID = msg.MessageID
		tgbotaddons.SendMsg(tgBot.bot, answerMsg)

	case "Нет":
		tgBot.chatsCache[msg.Chat.ID] = tgbotaddons.SetCache(mainMenu, nil)

		answerText := fmt.Sprintf("Очень жаль! Если всё таки захочешь завести, пиши /%s", createAnimal)
		answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
		answerMsg.ReplyToMessageID = msg.MessageID
		tgbotaddons.SendMsg(tgBot.bot, answerMsg)

	default:
		answerText := "Ну ответь на вопрос-то! Да или нет?"
		answerMsg := tgbotapi.NewMessage(msg.Chat.ID, answerText)
		answerMsg.ReplyToMessageID = msg.MessageID
		tgbotaddons.SendMsg(tgBot.bot, answerMsg)
	}
}
