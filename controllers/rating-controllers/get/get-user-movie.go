package getRating

import (
	"context"
	"log"
	"themoviebakery/config"
	"themoviebakery/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetRatingUserMovie(userId string, movieId string, mongoConnection *config.MongoConn) (*models.RatingTypeFullIdPrimitive, string) {
	statusCode := make(chan string, 1)

	var rating models.RatingTypeFullIdPrimitive
	err := mongoConnection.Collection.FindOne(context.TODO(), bson.M{"userid": userId, "movieid": movieId}).Decode(&rating)
	if err != nil {
		log.Println("Error trying to get rating by userid and movie =>", err)
		statusCode <- "ERROR_FINDING_RATING_BY_USER_404"
		return nil, <-statusCode
	} else {
		statusCode <- "nil"
	}

	return &rating, <-statusCode
}
