package createUser

import (
	"context"
	"net/http"
	config "themoviebakery/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(ginContext *gin.Context) {
	mongoNewConnection := config.ConnectMongo()

	res, err := mongoNewConnection.Collection.InsertOne(context.TODO(), bson.M{"yes": "we can"})
	if err != nil {
		panic(err)
	}

	defer mongoNewConnection.Disconnect()

	ginContext.IndentedJSON(http.StatusOK, res)
}
