package tg_bot

import (
	"animals-game/configs"
	"animals-game/internal/repository"
	tgbot_addons "animals-game/pkg/tgbot-addons"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

type TgBot struct {
	bot        *tgbotapi.BotAPI
	repo       *repository.Repository
	chatsCache map[int64]tgbot_addons.ChatCache
}

func NewTgBot(token string, repo *repository.Repository) TgBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	return TgBot{
		bot:        bot,
		repo:       repo,
		chatsCache: map[int64]tgbot_addons.ChatCache{},
	}
}

func (tgBot TgBot) Run() error {
	botConfig := configs.ParseBotConfig("./configs/tgbot-config.yaml")
	stickersConfig = configs.ParseStickersConfig("./configs/animals-stickers-config.yaml")

	updateConfig := tgbotapi.NewUpdate(botConfig.Offset)
	updateConfig.Timeout = botConfig.Timeout

	updates := tgBot.bot.GetUpdatesChan(updateConfig)

	slog.Info("Bot is ready to work!")

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
