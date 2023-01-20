package createUser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateRating(ginContext *gin.Context) {
	// mongoNewConnection := config.ConnectMongo()

	ginContext.IndentedJSON(http.StatusOK, bson.M{"fon": "boonha"})
}
