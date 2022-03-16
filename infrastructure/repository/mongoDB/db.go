package mongoDB

import (
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
)

// Init setting up mongDB
func Init() *mongo.Client {

	mongoURL := fmt.Sprintf("%s://%s:%s", os.Getenv("DB_TYPE"), os.Getenv("MONGO_DB_HOST"), os.Getenv("MONGO_DB_PORT"))

	mongoTimeout := time.Minute * 15

	// using go mongo-driver  to connect to mongoDB
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout))
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatalf("error %v", err)
	}

	log.Println("Database Connected Successfully...")
	return client
}
