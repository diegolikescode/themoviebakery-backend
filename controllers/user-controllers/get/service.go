package getUser

import (
	"context"
	"fmt"
	"net/http"
	"themoviebakery/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(ginContext *gin.Context) {
	queryParams := ginContext.Request.URL.Query()

	if queryParams["email"] == nil && queryParams["userId"] == nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, bson.M{"NO": "SHIT"})
		return
	}

	if queryParams["userId"] != nil {
		mongoNewConnection := config.ConnectMongo()
		defer mongoNewConnection.Disconnect()

		var result bson.D
		err := mongoNewConnection.Collection.FindOne(context.TODO(), bson.M{"userid": queryParams["userId"][0]}).Decode(&result)
		if err != nil {
			panic(err)
		}

		fmt.Println(result)
	}

	if queryParams["email"] != nil {
		mongoNewConnection := config.ConnectMongo()
		defer mongoNewConnection.Disconnect()

		filter := bson.M{"email": queryParams["email"]}

		cursor, err := mongoNewConnection.Collection.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}

		var results []bson.D
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}

		for _, result := range results {
			fmt.Print(result)
		}
	}
}
