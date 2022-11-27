package createUser

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserType struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email           string             `json:"email" validate:"required"`
	DisplayName     string             `json:"displayName" validate:"required"`
	Password        string             `json:"password" validate:"required"`
	ConfirmPassword string             `json:"confirmPassword" validate:"required"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}
