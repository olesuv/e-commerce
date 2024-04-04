package configs

import (
	"context"
	"fmt"
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

		if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
			panic(err)
		}

		fmt.Println("mongodb connected")
	})

	return client
}

type Collection interface {
	GetCollection(client *mongo.Client, collectionName string) *mongo.Collection

	FindOne(context.Context, interface{}, ...*options.FindOneOptions) *mongo.SingleResult
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(context.Context, []interface{}, ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(context.Context, interface{}, ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(context.Context, interface{}, ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := os.Getenv("dbName")
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}

func FindOne(collection Collection, ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return collection.FindOne(ctx, filter, opts...)
}

func InsertOne(collection Collection, ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(ctx, document, opts...)
}

func InsertMany(collection Collection, ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return collection.InsertMany(ctx, documents, opts...)
}

func UpdateOne(collection Collection, ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return collection.UpdateOne(ctx, filter, update, opts...)
}

func UpdateMany(collection Collection, ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return collection.UpdateMany(ctx, filter, update, opts...)
}

func DeleteOne(collection Collection, ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return collection.DeleteOne(ctx, filter, opts...)
}

func DeleteMany(collection Collection, ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return collection.DeleteMany(ctx, filter, opts...)
}
