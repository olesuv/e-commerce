package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type User struct {
	Id			primitive.ObjectID `bson:"_id"`
	Name 		string 						 `bson:"name,omitempty"`
	Email 	string 						 `bson:"email" validate:"required"` 
	Phone 	string 						 `bson:"phone,omitempty"`
	Orders	[]string					 `bson:"orders,omitempty"` 
}
