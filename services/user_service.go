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

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	_, err := us.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) DeleteUserById(id string) (*models.User, error) {
	user, err := us.GetUserById(id)
	if err != nil {
		return nil, err
	}
	_, err = us.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) DeleteUserByEmail(email string) (*models.User, error) {
	user, err := us.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	_, err = us.collection.DeleteOne(context.Background(), bson.M{"email": email})
	if err != nil {
		return nil, err
	}
	return user, nil
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

func (us *UserService) UpdateUserById(id string, user *models.User) (*models.User, error) {
	_, err := us.collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}
	updatedUser, err := us.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (us *UserService) UpdateUserByEmail(email string, user *models.User) (*models.User, error) {
	_, err := us.collection.UpdateOne(context.Background(), bson.M{"email": email}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}
	updatedUser, err := us.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
