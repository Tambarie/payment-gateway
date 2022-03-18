package payment_repository

import (
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository interface
type Repository interface {
	CheckIfEmailExists(email string) (bson.M, error)
	CheckIfUserExists(userReference string) (bson.M, error)
	CreateUser(user *domain.User) (*domain.User, error)
	Authorise(card *domain.Card) (*domain.Card, error)
	GetCardByID(id string) (bson.M, error)
	UpdateAccount(amount float64, id string) (interface{}, error)
	SaveCapturedTransaction(capture *domain.Transaction) (*mongo.InsertOneResult, error)
	GetCapturedTransactionByTransactionID(id string) (*domain.Transaction, error)
	GetCapturedTransactionByAuthorizationID(id string) (bson.M, error)
	RefundUpdateAccount(amount float64, id string, count int) (interface{}, error)
	SaveRefundTracker(tracker *domain.RefundTracker) (*mongo.InsertOneResult, error)
	GetRefundTrackerByTransactionID(id string) (bson.M, error)
	VoidCard(id string, void bool) (interface{}, error)
}
