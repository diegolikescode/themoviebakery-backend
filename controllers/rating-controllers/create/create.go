package createRating

import (
	"context"
	"net/http"
	"themoviebakery/config"
	getRating "themoviebakery/controllers/rating-controllers/get"
	updateRating "themoviebakery/controllers/rating-controllers/update"
	"themoviebakery/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func CreateRating(ginContext *gin.Context) {
	mongoNewConnection := config.ConnectMongo("ratings")

	var ratingBody models.RatingTypeFullIdPrimitive
	ginContext.ShouldBindJSON(&ratingBody)

	rating, statusCode := getRating.GetRatingUserMovie(ratingBody.UserId, ratingBody.MovieId, &mongoNewConnection)
	if statusCode != "nil" {
		res, statusCode := updateRating.UpdateUser(*rating, ratingBody, &mongoNewConnection)

		if statusCode == "nil" {
			ginContext.IndentedJSON(http.StatusOK, res)

			defer mongoNewConnection.Disconnect()
			return
		} else {
			ginContext.IndentedJSON(http.StatusInternalServerError, res)

			defer mongoNewConnection.Disconnect()
			return
		}
	}

	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		panic(err)
	}

	shortid.SetDefault(sid)

	shortId, err := sid.Generate()
	if err != nil {
		panic(err)
	}

	ratingBody.RatingId = shortId
	ratingBody.CreatedAt = time.Now()
	ratingBody.UpdatedAt = time.Now()

	res, err := mongoNewConnection.Collection.InsertOne(context.TODO(), ratingBody)
	if err != nil {
		panic(err)
	}

	defer mongoNewConnection.Disconnect()

	ginContext.IndentedJSON(http.StatusOK, res)
}
