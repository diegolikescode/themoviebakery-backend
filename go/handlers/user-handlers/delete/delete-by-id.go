package deleteUserHandler

import (
	"context"
	"log"
	"net/http"
	"themoviebakery/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUserById(ginContext *gin.Context) {
	queryParams := ginContext.Request.URL.Query()

	userObjectId, objectIdErr := primitive.ObjectIDFromHex(queryParams["userId"][0])
	if objectIdErr != nil {
		log.Println("Trying to parse user's object id (string) to mongo's ObjectId, Invalid ID =>", objectIdErr)
		ginContext.IndentedJSON(http.StatusOK, bson.M{"message": "couldn't transform this id to primitive.ObjectIDFromHex format"})
		return
	}

	mongoNewConnection := config.ConnectMongo("users")
	defer mongoNewConnection.Disconnect()

	res, err := mongoNewConnection.Collection.DeleteOne(context.TODO(), bson.M{"_id": userObjectId})

	if err == nil {
		if res.DeletedCount == 0 {
			ginContext.IndentedJSON(http.StatusNotFound, bson.M{"message": "user with id" + queryParams["userId"][0] + " wasn't founded"})
			return
		}
		ginContext.IndentedJSON(http.StatusOK, res)
	} else {
		ginContext.IndentedJSON(http.StatusInternalServerError, bson.M{"message": "user with id" + queryParams["userId"][0] + " wasn't founded"})
	}
}
