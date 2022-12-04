package updateUser

import (
	"context"
	"log"
	"net/http"
	config "themoviebakery/config"
	getUser "themoviebakery/controllers/user-controllers/get"
	models "themoviebakery/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(ginContext *gin.Context) {
	mongoNewConnection := config.ConnectMongo()
	var userBodyPrimitive models.UserTypeFullIdPrimitive
	ginContext.ShouldBindJSON(&userBodyPrimitive)

	currentUser, statusCode := getUser.GetUserById(userBodyPrimitive.Id.Hex(), &mongoNewConnection)
	if statusCode != "nil" {
		log.Println("user not founded by id")
		ginContext.IndentedJSON(http.StatusNotFound, bson.M{"message": "user not found by id"})
		return
	}

	otherUser, otherStatusCode := getUser.GetUserByEmail(userBodyPrimitive.Email, &mongoNewConnection)
	if otherStatusCode == "nil" {
		if currentUser.Id != otherUser.Id {
			log.Println("the new email is already taken")
			ginContext.IndentedJSON(http.StatusConflict, bson.M{"message": "the new email is already taken"})
			return
		}
	}

	userBodyPrimitive.UpdatedAt = time.Now()

	update := bson.D{{"$set", userBodyPrimitive}}

	res, err := mongoNewConnection.Collection.UpdateOne(context.TODO(), bson.M{"_id": userBodyPrimitive.Id}, update)
	if err != nil {
		panic(err)
	}

	defer mongoNewConnection.Disconnect()

	ginContext.IndentedJSON(http.StatusOK, res)
}
