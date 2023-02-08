package getRating

import (
	"context"
	"log"
	"net/http"
	"themoviebakery/config"
	"themoviebakery/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AllRatingsAllUsers(ginContext *gin.Context) {
	mongoConnection := config.ConnectMongo("ratings")
	defer mongoConnection.Disconnect()

	ratingCursor, err := mongoConnection.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error trying to get rating by userid =>", err)
		ginContext.IndentedJSON(http.StatusInternalServerError, bson.A{})
		return
	}

	var ratings []models.RatingMachineLearningData
	if err = ratingCursor.All(context.TODO(), &ratings); err != nil {
		panic(err)
	}

	ginContext.IndentedJSON(http.StatusOK, &ratings)
}
