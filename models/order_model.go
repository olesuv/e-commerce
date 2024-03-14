package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
    Id              primitive.ObjectID `bson:"_id"`
    Products        []string           `bson:"products,omitempty" validate:"required,minlen=1"`
    OrderDate       time.Time          `bson:"order_date,omitempty" validate:"required"`
    ShippingAddress string             `bson:"shipping_address,omitempty"`
    Status          string             `bson:"status,omitempty"`
    CustomerEmail   string             `bson:"email,omitempty" validate:"required"`
    PaymentStatus   string             `bson:"payment_status,omitempty"`
}

