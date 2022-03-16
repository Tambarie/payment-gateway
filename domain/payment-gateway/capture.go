package domain

type Capture struct {
	TransactionID   string  `json:"transaction_id" bson:"transaction_id"`
	AuthorizationID string  `json:"authorization_id" bson:"authorization_id"`
	Amount          float64 `json:"amount" bson:"amount"`
}
