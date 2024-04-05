package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server.go/configs"
	"server.go/models"
)

type UserService struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUserService() *UserService {
	client := configs.ConnectDB()
	collection := configs.GetCollection(client, "users")
	return &UserService{client, collection}
}

func (us *UserService) CreateUser(user *models.User) error {
	_, err := us.collection.InsertOne(context.Background(), user)
	return err
}

func (us *UserService) DeleteUserById(id string) error {
	_, err := us.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (us *UserService) DeleteUserByEmail(email string) error {
	_, err := us.collection.DeleteOne(context.Background(), bson.M{"email": email})
	return err
}

func (us *UserService) GetUserById(id string) (*models.User, error) {
	var user models.User
	err := us.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	return &user, err
}

func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := us.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (us *UserService) UpdateUserById(id string, user *models.User) error {
	_, err := us.collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": user})
	return err
}

func (us *UserService) UpdateUserByEmail(email string, user *models.User) error {
	_, err := us.collection.UpdateOne(context.Background(), bson.M{"email": email}, bson.M{"$set": user})
	return err
}
