package updateRating

import (
	"context"
	config "themoviebakery/config"
	"themoviebakery/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUser(oldRating models.RatingTypeFullIdPrimitive, updatedRating models.RatingTypeFullIdPrimitive, mongoConnection *config.MongoConn) (*mongo.UpdateResult, string) {
	statusCode := make(chan string, 1)

	var upMongoRating models.RatingTypeUpdate

	upMongoRating.UpdatedAt = time.Now()
	upMongoRating.RatingId = oldRating.RatingId
	upMongoRating.RatingValue = updatedRating.RatingValue
	upMongoRating.MovieId = oldRating.MovieId
	upMongoRating.UserId = oldRating.UserId

	update := bson.D{{"$set", upMongoRating}}
	res, err := mongoConnection.Collection.UpdateOne(context.TODO(), bson.M{"ratingid": oldRating.RatingId}, update)

	if err != nil {
		statusCode <- "ERROR_TRYING_TO_UPDATE_RATING"
		return nil, <-statusCode
	} else {
		statusCode <- "nil"
	}

	defer mongoConnection.Disconnect()
	return res, <-statusCode
}
