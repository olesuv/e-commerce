package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus int

const (
	Available OrderStatus = iota
	Buyed     OrderStatus = iota
)

type Order struct {
	Id              primitive.ObjectID `bson:"_id"`
	Title           []string           `bson:"order_title,omitempty" validate:"required,minlen=1"`
	Description     string             `bson:"description,omitempty"`
	Category        string             `bson:"category,omitempty"`
	Date            time.Time          `bson:"order_date,omitempty" validate:"required"`
	ShippingAddress string             `bson:"shipping_address,omitempty"`
	Status          OrderStatus        `bson:"status,omitempty"`
	CustomerEmail   string             `bson:"email,omitempty"`
}
