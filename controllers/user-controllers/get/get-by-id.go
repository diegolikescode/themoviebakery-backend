package getUser

import (
	"context"
	"log"
	"themoviebakery/config"
	createUser "themoviebakery/controllers/user-controllers/create"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserById(id string, mongoConnection *config.MongoConn) (*createUser.UserTypeFull, string) {
	statusCode := make(chan string, 1)

	userObjectId, objectIdErr := primitive.ObjectIDFromHex(id)
	if objectIdErr != nil {
		log.Println("Trying to parse user's object id (string) to mongo's ObjectId, Invalid ID =>", objectIdErr)
	}
	var user createUser.UserTypeFull
	err := mongoConnection.Collection.FindOne(context.TODO(), bson.M{"_id": userObjectId}).Decode(&user)
	if err != nil {
		log.Println("Trying to parse user's object id (string) to mongo's ObjectId, Invalid ID =>", err)
		statusCode <- "ERROR_FINDING_USER_BY_ID_404"
		return nil, <-statusCode
	} else {
		statusCode <- "nil"
	}

	return &user, <-statusCode
}
