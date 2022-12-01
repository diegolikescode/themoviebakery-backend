package updateUser

import (
	"context"
	"net/http"
	config "themoviebakery/config"
	createUser "themoviebakery/controllers/user-controllers/create"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(ginContext *gin.Context) {
	mongoNewConnection := config.ConnectMongo()

	var userBody createUser.UserTypeFull
	ginContext.ShouldBindJSON(&userBody)

	userBody.UpdatedAt = time.Now()

	update := bson.D{{"$set", userBody}}

	// res, err := mongoNewConnection.Collection.UpdateOne(context.TODO(), bson.M{"_id": userBody.Id}, {$set: userBody})
	res, err := mongoNewConnection.Collection.UpdateOne(context.TODO(), bson.M{"_id": userBody.Id}, update)
	if err != nil {
		panic(err)
	}

	defer mongoNewConnection.Disconnect()

	ginContext.IndentedJSON(http.StatusOK, res)
}
