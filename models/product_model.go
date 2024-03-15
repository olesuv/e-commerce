package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Title       string             `bson:"title" validate:"required,maxLength=100"`
    Description string             `bson:"description" validate:"required,maxLength=200"`
    Amount      int                `bson:"amount" validate:"required"`
    Price       float64            `bson:"price" validate:"required"`
    Images      []string           `bson:"images,omitempty"`
    Category    string             `bson:"category" validate:"required"`
}

