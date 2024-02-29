package tg_bot

import (
	bot_configs "animals-game/bot-configs"
	"animals-game/configs"
)

var botConfig *bot_configs.BotConfig
var stickersConfig *bot_configs.StickersConfig
var animalTypeFromIdToStr configs.AnimalTypeFromIdToStr
var animalTypeFromStrToId configs.AnimalTypeFromStrToId

func initConfigs() {
	botConfig = bot_configs.ParseBotConfig("./bot-configs/tgbot-config.yaml")
	stickersConfig = bot_configs.ParseStickersConfig("./bot-configs/animals-stickers-config.yaml")
	animalsCfg := configs.ParseAnimalsType("./configs/animals-types-config.yaml")
	animalTypeFromIdToStr = configs.GetAnimalsTYpeListFromIdToStr(animalsCfg)
	animalTypeFromStrToId = configs.GetAnimalsTypeListFromStrToId(&animalTypeFromIdToStr)
}
