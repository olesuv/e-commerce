package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server.go/configs"
	"server.go/models"
)

type UserAuthService struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUserVerficicationService() *UserAuthService {
	client := configs.ConnectDB()
	collection := configs.GetCollection(client, "sessions")
	return &UserAuthService{client, collection}
}

func (uvs *UserAuthService) CreateSession(session *models.Session) (*models.Session, error) {
	_, err := uvs.collection.InsertOne(context.Background(), session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (uvs *UserAuthService) DeleteSession(email string) (*models.Session, error) {
	session, err := uvs.GetSessionByEmail(email)
	if err != nil {
		return nil, err
	}
	_, err = uvs.collection.DeleteOne(context.Background(), bson.M{"user_email": email})
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (uvs *UserAuthService) UpdateSession(email string, session *models.Session) (*models.Session, error) {
	_, err := uvs.collection.ReplaceOne(context.Background(), bson.M{"user_email": email}, session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (uvs *UserAuthService) GetSessionByToken(token string) (*models.Session, error) {
	var session models.Session
	err := uvs.collection.FindOne(context.Background(), bson.M{"token": token}).Decode(&session)
	return &session, err
}

func (uvs *UserAuthService) GetSessionByEmail(email string) (*models.Session, error) {
	var session models.Session
	err := uvs.collection.FindOne(context.Background(), bson.M{"user_email": email}).Decode(&session)
	return &session, err
}

func (uvs *UserAuthService) GetSessionById(id string) (*models.Session, error) {
	var session models.Session
	err := uvs.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&session)
	return &session, err
}
