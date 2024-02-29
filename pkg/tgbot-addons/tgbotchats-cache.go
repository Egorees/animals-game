package tgbot_addons

import "time"

type StatusType int16

type ChatCache struct {
	Command         StatusType
	InteractionTime time.Time
	OtherInfo       interface{}
}

func SetCache(cmd StatusType, otherInfo interface{}) ChatCache {
	return ChatCache{
		Command:         cmd,
		InteractionTime: time.Now(),
		OtherInfo:       otherInfo,
	}
}
