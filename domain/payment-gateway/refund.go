package domain

type Refund struct {
	AuthorizationID string  `json:"authorization_id" bson:"authorization_id"`
	Amount          float64 `json:"amount" bson:"amount"`
}
