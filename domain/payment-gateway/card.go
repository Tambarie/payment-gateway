package domain

type Card struct {
	UserReference   string  `json:"user_reference" bson:"user_reference"`
	ID              string  `json:"id,omitempty" bson:"id"`
	CardNumber      int64   `json:"card_number" bson:"card_number"`
	ExpirationYear  string  `json:"expiration_year" bson:"expiration_year"`
	ExpirationMonth string  `json:"expiration_month" bson:"expiration_month"`
	Cvv             int64   `json:"cvv" bson:"cvv"`
	Amount          float64 `json:"amount" bson:"amount"`
	Currency        string  `json:"currency" bson:"currency"`
	Count           int     `json:"count" bson:"count"`
	Void            bool    `json:"-" bson:"void"`
}

func (c *Card) VoidCard() {
	c.Void = true
}
