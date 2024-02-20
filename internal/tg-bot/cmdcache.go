package tg_bot

import "time"

// users commands ids
const (
	MainMenu    = 1
	SetUsername = 2
)

// chats commands id's
const ()

type cmdCache struct {
	lastCmd int16
	cmdTime time.Time
}
