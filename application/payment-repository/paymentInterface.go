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
	UpdateAccount(amount float64, id string) (*mongo.UpdateResult, error)
	SaveCapturedTransaction(capture *domain.Capture) (*domain.Capture, error)
}
