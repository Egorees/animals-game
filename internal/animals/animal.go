package animals

type Animal struct {
	Id          int   `json:"Id"`
	Type        int16 `json:"Type"`
	Exp         int64 `json:"Exp"`
	OwnerUserId int   `json:"OwnerUserId"`
}
