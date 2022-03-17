package payment_repository

import (
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository interface
type Repository interface {
	CreateMerchant(card *domain.Card) (*domain.Card, error)
	GetID(id string) (bson.M, error)
	UpdateAccount(amount float64, id string) (interface{}, error)
	SaveCapturedTransaction(capture *domain.Capture) (*mongo.InsertOneResult, error)
	GetCapturedTransactionByTransactionID(id string) (*domain.Capture, error)
	GetCapturedTransactionByAuthorizationID(id string) (bson.M, error)
	RefundUpdateAccount(amount float64, id string, count int) (interface{}, error)
	SaveRefundTracker(tracker *domain.RefundTracker) (*mongo.InsertOneResult, error)
	GetRefundTrackerByTransactionID(id string) (bson.M, error)
}
