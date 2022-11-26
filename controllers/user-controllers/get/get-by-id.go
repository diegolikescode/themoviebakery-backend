package getUser

import (
	"context"
	"log"
	"net/http"
	"themoviebakery/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserById(ginContext *gin.Context) {
	queryParams := ginContext.Request.URL.Query()

	mongoNewConnection := config.ConnectMongo()
	defer mongoNewConnection.Disconnect()

	userObjectId, objectIdErr := primitive.ObjectIDFromHex(queryParams["userId"][0])
	if objectIdErr != nil {
		log.Println("Trying to parse user's object id (string) to mongo's ObjectId, Invalid ID =>", objectIdErr)
	}

	var result bson.D
	err := mongoNewConnection.Collection.FindOne(context.TODO(), bson.M{"_id": userObjectId}).Decode(&result)
	if err != nil {
		log.Println("Trying to parse user's object id (string) to mongo's ObjectId, Invalid ID =>", err)
		ginContext.IndentedJSON(http.StatusInternalServerError, bson.M{"message": "problem trying to get user's profile using ObjectId"})
		return
	}

	ginContext.IndentedJSON(http.StatusOK, &result)
}
