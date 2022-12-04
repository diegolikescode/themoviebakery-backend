package getUser

import (
	"context"
	"log"
	"themoviebakery/config"
	models "themoviebakery/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByEmail(email string, mongoConnection *config.MongoConn) (*models.UserTypeFull, string) {
	statusCode := make(chan string, 1)

	var user models.UserTypeFull
	err := mongoConnection.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Println("error trying to get the user by it's email =>", err)
		statusCode <- "USER_NOT_FOUND_BY_EMAIL_404"
	} else {
		statusCode <- "nil"
	}

	return &user, <-statusCode
}
