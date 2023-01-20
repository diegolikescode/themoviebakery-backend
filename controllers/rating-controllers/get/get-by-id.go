package getRating

import (
	"context"
	"log"
	"themoviebakery/config"
	"themoviebakery/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetRatingById(id string, mongoConnection *config.MongoConn) (*models.RatingTypeFullIdPrimitive, string) {
	statusCode := make(chan string, 1)

	var rating models.RatingTypeFullIdPrimitive
	err := mongoConnection.Collection.FindOne(context.TODO(), bson.M{"ratingid": id}).Decode(&rating)
	if err != nil {
		log.Println("Error trying to get rating by ratingId =>", err)
		statusCode <- "ERROR_FINDING_RATING_BY_ID_404"
		return nil, <-statusCode
	} else {
		statusCode <- "nil"
	}

	return &rating, <-statusCode
}
