package updateUser

import (
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
	postgresConnection := config.ConnectPostgres()
	var userBodyPrimitive models.UserTypeFullIdPrimitive
	ginContext.ShouldBindJSON(&userBodyPrimitive)

	currentUser, statusCode := getUser.GetUserById(userBodyPrimitive.Id.Hex(), &postgresConnection)
	if statusCode != "nil" {
		log.Println("user not founded by id")
		ginContext.IndentedJSON(http.StatusNotFound, bson.M{"message": "user not found by id"})
		return
	}

	otherUser, otherStatusCode := getUser.GetUserByEmail(userBodyPrimitive.Email, &postgresConnection)
	if otherStatusCode == "nil" {
		if currentUser.Id != otherUser.Id {
			log.Println("the new email is already taken")
			ginContext.IndentedJSON(http.StatusConflict, bson.M{"message": "the new email is already taken"})
			return
		}
	}

	userBodyPrimitive.UpdatedAt = time.Now()

	res, err := postgresConnection.DbConn.Exec("")
	if err != nil {
		panic(err)
	}

	defer postgresConnection.Disconnect()

	ginContext.IndentedJSON(http.StatusOK, res)
}
