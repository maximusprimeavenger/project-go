package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	User_name      *string            `json:"user_name" validate:"required,min=5,max=100"`
	User_last_name *string            `json:"user_last_name" validate:"required,min=5,max=100"`
	User_password  *string            `json:"password" validate:"required,min=5"`
	User_email     *string            `json:"email" validate:"email,required"`
	User_phone     *string            `json:"phone" validate:"required"`
	Token          *string            `json:"token"`
	User_type      *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token  *string            `json:"refresh_token"`
	Created_at     time.Time          `json:"created_at"`
	Updated_at     time.Time          `json:"upadated_at"`
	User_id        string             `json:"user_id"`
}
