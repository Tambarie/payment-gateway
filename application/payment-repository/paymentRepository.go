package payment_repository

import (
	"fmt"
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

// RepositoryDB struct
type RepositoryDB struct {
	db *mongo.Client
}

// NewWalletRepositoryDB function to initialize RepositoryDB struct
func NewPaymentGatewayRepositoryDB(client *mongo.Client) *RepositoryDB {
	return &RepositoryDB{
		db: client,
	}
}

func (paymentRepo *RepositoryDB) CreateMerchant(card *domain.Card) (*domain.Card, error) {
	collection := paymentRepo.db.Database("payment-gateway").Collection("gateway")
	result, err := collection.InsertOne(context.TODO(), card)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return card, err
}

// CreateWallet  Creates payment-gateway to the database

// CheckIfPasswordExists   Check if the user password exists in the database

// SaveTransaction Saving the payment-gateway transaction to the database

// PostToAccount Posting the account details to the database

// GetAccountBalance Getting account balance from the database

// ChangeUserStatus Updating user status from the database
