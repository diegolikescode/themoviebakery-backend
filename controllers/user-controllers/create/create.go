package createUser

import (
	"context"
	"fmt"
	"net/http"
	config "themoviebakery/config"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(ginContext *gin.Context) {
	mongoNewConnection := config.ConnectMongo()

	var userBody UserTypeInsert
	ginContext.ShouldBindJSON(&userBody)
	fmt.Println("XAMA", userBody)
	userBody.CreatedAt = time.Now()
	userBody.UpdatedAt = time.Now()
	fmt.Println("XAMA DE NOVO", userBody)

	res, err := mongoNewConnection.Collection.InsertOne(context.TODO(), userBody)
	if err != nil {
		panic(err)
	}

	defer mongoNewConnection.Disconnect()

	ginContext.IndentedJSON(http.StatusOK, res)
}
