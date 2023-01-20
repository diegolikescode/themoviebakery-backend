package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingTypeInsert struct {
	RatingId    string    `json:"ratingId"`
	UserId      string    `json:"userId"`
	MovieId     string    `json:"movieId"`
	RatingValue float32   `json:"ratingValue"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RatingTypeFullIdPrimitive struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	RatingId    string             `json:"ratingId"`
	UserId      string             `json:"userId"`
	MovieId     string             `json:"movieId"`
	RatingValue float32            `json:"ratingValue"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type RatingTypeUpdate struct {
	RatingId    string    `json:"ratingId"`
	UserId      string    `json:"userId"`
	MovieId     string    `json:"movieId"`
	RatingValue float32   `json:"ratingValue"`
	UpdatedAt   time.Time `json:"updated_at"`
}
