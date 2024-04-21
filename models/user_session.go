package models

type Session struct {
	Id          string `bson:"_id"`
	UserEmail   string `bson:"user_email" validate:"required"`
	DateExpired string `bson:"date_expired" validate:"required"`
	Token       string `bson:"token" validate:"required"`
}
