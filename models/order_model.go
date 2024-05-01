package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus int
type OrderCategory int

const (
	Available OrderStatus = iota
	Buyed
)

const (
	Electronics OrderCategory = iota
	Fashion
	Home
	Sports
	Books
	Automotive
	Other
)

type Order struct {
	Id              primitive.ObjectID `bson:"_id"`
	Title           string             `bson:"order_title,omitempty" validate:"required,minlen=3,maxlen=100"`
	Description     string             `bson:"description,omitempty" validate:"required,minlen=3,maxlen=1000"`
	Images          []primitive.Binary `bson:"images,omitempty" validate:"required"`
	Category        OrderCategory      `bson:"category,omitempty"`
	Date            time.Time          `bson:"order_date,omitempty" validate:"required"`
	ShippingAddress string             `bson:"shipping_address,omitempty"`
	Status          OrderStatus        `bson:"status,omitempty"`
	CustomerEmail   string             `bson:"email,omitempty"`
}
