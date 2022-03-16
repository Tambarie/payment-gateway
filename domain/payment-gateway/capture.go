package domain

type Capture struct {
	AuthorizationID string  `json:"authorization_id" bson:"authorization_id"`
	Amount          float64 `json:"amount" bson:"amount"`
}
