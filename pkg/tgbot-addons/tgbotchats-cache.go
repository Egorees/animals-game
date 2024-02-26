package tgbot_addons

import "time"

type CommandType int16

type ChatCache struct {
	Command         CommandType
	InteractionTime time.Time
	OtherInfo       interface{}
}

func SetCache(cmd CommandType, interactionTime time.Time, otherInfo interface{}) ChatCache {
	return ChatCache{
		Command:         cmd,
		InteractionTime: interactionTime,
		OtherInfo:       otherInfo,
	}
}
