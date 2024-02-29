package tg_bot

import tgbot_addons "animals-game/pkg/tgbot-addons"

const (
	mainMenu = tgbot_addons.StatusType(iota)
	waitAnimalType
	waitAnimalName
	waitAcceptAnimalInfo
)
