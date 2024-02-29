package animals

type Animal struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Type    int16  `json:"Type"`
	Exp     int64  `json:"Exp"`
	OwnerId int    `json:"OwnerId"`
}

func NewAnimal() *Animal {
	return &Animal{
		Exp: 0,
	}
}
