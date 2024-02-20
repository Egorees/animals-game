package tg_bot

import (
	"animals-game/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type TgBot struct {
	bot        *tgbotapi.BotAPI
	repo       *repository.Repository
	chatsCache *map[string]cmdCache
}

func NewTgBot(token string, repo *repository.Repository) TgBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	return TgBot{
		bot:        bot,
		repo:       repo,
		chatsCache: &map[string]cmdCache{},
	}
}

func (tgBot TgBot) Run() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	err := tgBot.repo.CreateUserWithTgId(1)
	if err != nil {
		return err
	}

	if err != nil {
		log.Fatal(err)
	}

	updates := tgBot.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		switch {
		case update.FromChat().IsPrivate():
			privateHandler(tgBot, &update)
		case update.FromChat().IsPrivate() || update.FromChat().IsSuperGroup():
			groupHandler(tgBot, &update)
		default:
			continue
		}
	}
	return nil
}
