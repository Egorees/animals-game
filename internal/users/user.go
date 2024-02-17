package users

type User struct {
	Id         int64  `json:"Id"`
	TelegramId string `json:"TelegramId"`
	Login      string `json:"Login"`
	Password   string `json:"Password"`
	AnimalId   int64  `json:"AnimalId"`
}
