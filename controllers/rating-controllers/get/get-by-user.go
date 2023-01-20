package getRating

import (
	"context"
	"log"
	"themoviebakery/config"
	"themoviebakery/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetRatingByUser(id string, mongoConnection *config.MongoConn) (*[]models.RatingTypeFullIdPrimitive, string) {
	statusCode := make(chan string, 1)

	ratingCursor, err := mongoConnection.Collection.Find(context.TODO(), bson.M{"userid": id})
	if err != nil {
		log.Println("Error trying to get rating by userid =>", err)
		statusCode <- "ERROR_FINDING_RATING_BY_USER_404"
		return nil, <-statusCode
	} else {
		statusCode <- "nil"
	}

	var userRatings []models.RatingTypeFullIdPrimitive
	if err = ratingCursor.All(context.TODO(), &userRatings); err != nil {
		panic(err)
	}

	return &userRatings, <-statusCode
}
