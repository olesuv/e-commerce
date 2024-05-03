package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server.go/configs"
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
	var order models.Order
	err := os.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&order)
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
