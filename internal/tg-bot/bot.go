package tg_bot

import (
	"animals-game/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
)

type TgBot struct {
	bot  *tgbotapi.BotAPI
	repo *repository.Repository
}

func NewTgBot(token string) TgBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	return TgBot{bot: bot}
}

func (tgBot TgBot) RepoInit(db *sqlx.DB) {
	tgBot.repo = repository.NewRepository(db)
}

func (tgBot TgBot) Run() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := tgBot.bot.GetUpdatesChan(updateConfig)

	for update := range updates {

		if update.Message == nil { //later will add others
			continue
		}

		switch {
		case update.Message.Chat.Type == "private":
			tgBot.PrivateMessageHandler(update.Message)
		case update.Message.Chat.Type == "group" || update.Message.Chat.Type == "supergroup":
			tgBot.GroupMessageHandler(update.Message)
		default:
			continue
		}
	}
	return nil
}
