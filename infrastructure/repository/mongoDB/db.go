package mongoDB

import (
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

	//mongoURL := fmt.Sprintf("%s://%s:%s", os.Getenv("DB_TYPE"), os.Getenv("MONGO_DB_HOST"), os.Getenv("MONGO_DB_PORT"))
	//
	//mongoTimeout := time.Minute * 15
	//
	//// using go mongo-driver  to connect to mongoDB
	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout))
	//defer cancel()
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	//if err != nil {
	//	log.Fatalf("error %v", err)
	//}
	//
	//log.Println("Database Connected Successfully...")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGODB_URI")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connected Successfully...")
	return client
}
