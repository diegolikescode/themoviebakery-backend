package getRatingHandler

import (
	"net/http"
	"themoviebakery/config"
	getRating "themoviebakery/controllers/rating-controllers/get"

	"github.com/gin-gonic/gin"
)

func GetRatingUserMovie(ginContext *gin.Context) {
	queryParams := ginContext.Request.URL.Query()

	mongoNewConnection := config.ConnectMongo("ratings")
	defer mongoNewConnection.Disconnect()
	rating, statusCode := getRating.GetRatingUserMovie(queryParams["userId"][0], queryParams["movieId"][0], &mongoNewConnection)

	if statusCode == "nil" {
		ginContext.IndentedJSON(http.StatusOK, &rating)
	} else {
		ginContext.IndentedJSON(http.StatusNotFound, &rating)
	}
}
