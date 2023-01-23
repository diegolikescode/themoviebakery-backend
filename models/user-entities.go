package models

import (
	"time"

	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserTypeInsert struct {
	UserId      string    `json:"userId"`
	Email       string    `json:"email" validate:"required"`
	DisplayName string    `json:"displayName" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserTypeFullIdPrimitive struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId      shortid.Shortid    `json:"userId"`
	Email       string             `json:"email" validate:"required"`
	DisplayName string             `json:"displayName" validate:"required"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type UserEssentialData struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId      string             `json:"userId"`
	Email       string             `json:"email"`
	DisplayName string             `json:"displayName"`
}
