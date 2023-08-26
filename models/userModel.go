package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    string             `json:"first_name" validate:"required,min=2,max=100"`
	LastName     string             `json:"last_name" validate:"required,min=2,max=100"`
	Password     string             `json:"password" validate:"required,min=6,max=20"`
	Email        string             `json:"email" validate:"required"`
	Phone        string             `json:"phone" validate:"required"`
	AccessToken  string             `json:"token"`
	UserType     string             `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken string             `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserID       string             `json:"user_id"`
}
