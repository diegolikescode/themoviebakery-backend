package getUserHandler

import (
	"net/http"
	"themoviebakery/config"
	getUser "themoviebakery/controllers/user-controllers/get"

	"github.com/gin-gonic/gin"
)

func HandlerGetUserByEmail(ginContext *gin.Context) {
	queryParams := ginContext.Request.URL.Query()

	mongoNewConnection := config.ConnectMongo()
	defer mongoNewConnection.Disconnect()

	user, statusCode := getUser.GetUserByEmail(queryParams["email"][0], &mongoNewConnection)

	if statusCode == "nil" {
		ginContext.IndentedJSON(http.StatusOK, &user)
	} else {
		ginContext.IndentedJSON(http.StatusNotFound, &user)
	}

}
