package users

type User struct {
	Id         int    `json:"Id"`
	TelegramId string `json:"TelegramId"`
	Login      string `json:"Login"`
	Password   string `json:"Password"`
	AnimalId   int    `json:"AnimalId"`
}
