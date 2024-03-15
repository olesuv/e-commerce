package configs

import (
  "context"
  "fmt"
	"log"
  "os"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"

	"github.com/lpernett/godotenv"
)

func ConnectDB() *mongo.Client{
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  connectionString:= os.Getenv("connectionString")

  // Use the SetServerAPIOptions() method to set the version of the Stable API on the client
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

  // Create a new client and connect to the server
  client, err := mongo.Connect(context.TODO(), opts)
  if err != nil {
    panic(err)
  }

  defer func() {
    if err = client.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()

  // Send a ping to confirm a successful connection
  if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
    panic(err)
  }

  fmt.Println("mongodb connected")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := os.Getenv("dbName")
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}

