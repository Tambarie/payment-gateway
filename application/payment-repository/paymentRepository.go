package payment_repository

import (
	"fmt"
	"github.com/Tambarie/payment-gateway/domain/helpers"
	domain "github.com/Tambarie/payment-gateway/domain/payment-gateway"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
)

// RepositoryDB struct
type RepositoryDB struct {
	db *mongo.Client
}

// NewPaymentGatewayRepositoryDB function to initialize RepositoryDB struct
func NewPaymentGatewayRepositoryDB(client *mongo.Client) *RepositoryDB {
	return &RepositoryDB{
		db: client,
	}
}

func (paymentRepo *RepositoryDB) CreateMerchant(card *domain.Card) (*domain.Card, error) {

	helpers.LogEvent("INFO", fmt.Sprintf("Creating Merchant with reference: %s ...", card))

	collection := paymentRepo.db.Database("payment-gateway").Collection("gateway")
	result, err := collection.InsertOne(context.TODO(), card)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return card, err
}

func (paymentRepo *RepositoryDB) GetCardByID(id string) (bson.M, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Getting card by ID  with id: %s ...", id))
	collection := paymentRepo.db.Database("payment-gateway").Collection("gateway")

	var result bson.M
	err := collection.FindOne(
		context.TODO(),
		bson.D{{"id", id}},
	).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
	}
	fmt.Printf("found document %v", result)
	return result, nil
}

func (paymentRepo *RepositoryDB) UpdateAccount(amount float64, id string) (interface{}, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Updating account with amount: %v and id :%s", amount, id))
	collection := paymentRepo.db.Database("payment-gateway").Collection("gateway")

	filter := bson.D{{"id", id}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{{"$set", bson.D{{"amount", amount}}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")

	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
	return result.UpsertedID, err
}

func (paymentRepo *RepositoryDB) SaveCapturedTransaction(capture *domain.Transaction) (*mongo.InsertOneResult, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Saving Captured transaction with reference: %s ...", capture))
	collection := paymentRepo.db.Database("payment-gateway").Collection("captured-transaction")
	result, err := collection.InsertOne(context.TODO(), capture)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return result, err
}

func (paymentRepo *RepositoryDB) GetCapturedTransactionByTransactionID(id string) (*domain.Transaction, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Getting Captured transaction by transaction ID :%s", id))
	collection := paymentRepo.db.Database("payment-gateway").Collection("captured-transaction")
	var result *domain.Transaction
	err := collection.FindOne(
		context.TODO(),
		bson.D{{"transaction_id", id}},
	).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
	}
	fmt.Printf("found document in captured-transaction collection %v", result)
	return result, nil
}

func (paymentRepo *RepositoryDB) GetCapturedTransactionByAuthorizationID(id string) (bson.M, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Getting Captured transaction by authorization ID :%s", id))
	collection := paymentRepo.db.Database("payment-gateway").Collection("captured-transaction")
	var result bson.M
	err := collection.FindOne(
		context.TODO(),
		bson.D{{"authorization_id", id}},
	).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
	}
	fmt.Printf("found document in captured-transaction collection %v", result)
	return result, nil
}

func (paymentRepo *RepositoryDB) RefundUpdateAccount(amount float64, id string, count int) (interface{}, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Updating account with :%v and id :%s ", amount, id))
	collection := paymentRepo.db.Database("payment-gateway").Collection("gateway")

	filter := bson.D{{"id", id}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{{"$set", bson.D{{"amount", amount}, {"count", count}}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")

	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
	return result.UpsertedID, err
}

func (paymentRepo *RepositoryDB) SaveRefundTracker(tracker *domain.RefundTracker) (*mongo.InsertOneResult, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Saving refund tracker with  :%v", tracker))
	collection := paymentRepo.db.Database("payment-gateway").Collection("refund-tracker-collection")
	result, err := collection.InsertOne(context.TODO(), tracker)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return result, err
}

func (paymentRepo *RepositoryDB) GetRefundTrackerByTransactionID(id string) (bson.M, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Getting Refund tracker by transaction ID :%s", id))
	collection := paymentRepo.db.Database("payment-gateway").Collection("refund-tracker-collection")
	var result bson.M
	err := collection.FindOne(
		context.TODO(),
		bson.D{{"transaction_id", id}},
	).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
	}
	fmt.Printf("found document in refund-tracker-collection %v", result)
	return result, nil
}

func (paymentRepo *RepositoryDB) VoidCard(id string, void bool) (interface{}, error) {

	helpers.LogEvent("INFO", fmt.Sprintf("Void card with id of  :%s", id))

	collection := paymentRepo.db.Database("payment-gateway").Collection("gateway")

	filter := bson.D{{"id", id}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{{"$set", bson.D{{"void", void}}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")

	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
	return result.UpsertedID, err
}
