package createUser

import (
	"context"
	"net/http"
	config "themoviebakery/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(ginContext *gin.Context) {
	mongoNewConnection := config.ConnectMongo()

	var userBody InputCreateUser
	ginContext.ShouldBindJSON(&userBody)

	newId := uuid.New()
	userBody.UserID = newId.String()

	res, err := mongoNewConnection.Collection.InsertOne(context.TODO(), userBody)
	if err != nil {
		panic(err)
	}

	defer mongoNewConnection.Disconnect()

	ginContext.IndentedJSON(http.StatusOK, res)
}
