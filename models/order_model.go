package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"server.go/constants"
)

type Order struct {
	Id              primitive.ObjectID        `bson:"_id"`
	Title           string                    `bson:"order_title,omitempty" validate:"required,minlen=3,maxlen=100"`
	Description     string                    `bson:"description,omitempty" validate:"required,minlen=3,maxlen=1000"`
	Images          []primitive.Binary        `bson:"images,omitempty" validate:"required"`
	Category        constants.OrderCategories `bson:"category,omitempty"`
	Date            time.Time                 `bson:"order_date,omitempty" validate:"required"`
	ShippingAddress string                    `bson:"shipping_address,omitempty"`
	Status          constants.OrderStatus     `bson:"status,omitempty"`
	AuthorEmail     string                    `bson:"author_email" validate:"required"`
	CustomerEmail   string                    `bson:"email,omitempty"`
	Price           float64                   `bson:"price,omitempty" validate:"required"`
	Currency        constants.OrderCurrency   `bson:"currency,omitempty" validate:"required"`
}
