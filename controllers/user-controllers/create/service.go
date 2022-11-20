package createUser

import (
	"fmt"
	"net/http"
	config "themoviebakery/config"

	"github.com/gin-gonic/gin"
)

func CreateUser(ginContext *gin.Context) {
	var mongoNewConnection = config.ConnectMongo()
	fmt.Println(mongoNewConnection)

	// res, err := mongoNewConnection.collection.InsertOne(context.TODO(), bson.M{"hello": "bebie"})
	// if err != nil {
	// 	panic(err)
	// }

	ginContext.IndentedJSON(http.StatusOK, mongoNewConnection)

	// return id
}
