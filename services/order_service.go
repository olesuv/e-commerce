package services

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server.go/configs"
	"server.go/constants"
	"server.go/models"
)

type OrderService struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewOrderService() *OrderService {
	client := configs.ConnectDB()
	collection := configs.GetCollection(client, "orders")
	return &OrderService{client, collection}
}

func (os *OrderService) CreateOrder(order *models.Order) (*models.Order, error) {
	_, err := os.collection.InsertOne(context.Background(), order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (os *OrderService) GetOrderById(id string) (*models.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var order models.Order
	err = os.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (os *OrderService) DeleteOrderById(id string) (*models.Order, error) {
	order, err := os.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	_, err = os.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (os *OrderService) GetOrders() ([]models.Order, error) {
	var orders []models.Order

	cursor, err := os.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var order models.Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}

	return orders, nil
}

func (os *OrderService) UpdateOrder(order *models.Order) (*models.Order, error) {
	_, err := os.collection.UpdateOne(context.Background(), bson.M{"_id": order.Id}, bson.M{"$set": order})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (os *OrderService) Last5Orders() ([]models.Order, error) {
	var orders []models.Order

	filter := bson.M{"status": constants.Available}

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"order_date": -1})
	findOptions.SetLimit(5)

	cursor, err := os.collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("server: get latest orders, details: %w", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var order models.Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}

	return orders, nil
}
