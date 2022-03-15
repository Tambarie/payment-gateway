package paymentGatewayRepository

import (
	"go.mongodb.org/mongo-driver/mongo"
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

// GetUserByEmail Fetches user based on the email from the database

// CreateWallet  Creates payment-gateway to the database

// CheckIfPasswordExists   Check if the user password exists in the database

// SaveTransaction Saving the payment-gateway transaction to the database

// PostToAccount Posting the account details to the database

// GetAccountBalance Getting account balance from the database

// ChangeUserStatus Updating user status from the database
