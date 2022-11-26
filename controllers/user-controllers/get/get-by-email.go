package getUser

import (
	"context"
	"log"
	"net/http"
	"themoviebakery/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByEmail(ginContext *gin.Context) {
	queryParams := ginContext.Request.URL.Query()

	mongoNewConnection := config.ConnectMongo()
	defer mongoNewConnection.Disconnect()

	var result bson.D
	err := mongoNewConnection.Collection.FindOne(context.TODO(), bson.M{"email": queryParams["email"][0]}).Decode(&result)
	if err != nil {
		log.Println("error trying to get the user by it's email =>", err)
		ginContext.IndentedJSON(http.StatusInternalServerError, bson.M{"message": "probably wrong params, is expected: userId and email"})
		return
	}

	ginContext.IndentedJSON(http.StatusOK, result)

}
