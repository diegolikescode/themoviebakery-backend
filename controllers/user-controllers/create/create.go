package createUser

import (
	"context"
	"net/http"
	config "themoviebakery/config"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(ginContext *gin.Context) {
	mongoNewConnection := config.ConnectMongo()

	var userBody UserType
	ginContext.ShouldBindJSON(&userBody)

	userBody.CreatedAt = time.Now()
	userBody.UpdatedAt = time.Now()

	res, err := mongoNewConnection.Collection.InsertOne(context.TODO(), userBody)
	if err != nil {
		panic(err)
	}

	defer mongoNewConnection.Disconnect()

	ginContext.IndentedJSON(http.StatusOK, res)
}
