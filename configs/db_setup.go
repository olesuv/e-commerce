package configs

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client           *mongo.Client
	once             sync.Once
	connectionString string
	err              error
)

func ConnectDB() *mongo.Client {
	once.Do(func() {
		connectionString = os.Getenv("connectionString")

		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

		client, err = mongo.Connect(context.TODO(), opts)
		if err != nil {
			panic(err)
		}

		if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
			panic(err)
		}

		log.Printf("connected to mongodb")
	})

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := os.Getenv("dbName")
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
