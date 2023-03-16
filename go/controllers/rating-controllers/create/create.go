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

	// PRINT JSON
	// var buffer bytes.Buffer
	// enc := json.NewEncoder(&buffer)
	// enc.SetIndent("", "  ")
	// if err := enc.Encode(ratingBody); err != nil {
	// 	panic("encode string err!")
	// }
	// fmt.Println(buffer.String())
	// END PRINT JSON

	rating, statusCode := getRating.GetRatingUserMovie(ratingBody.UserId, ratingBody.MovieId, &mongoNewConnection)
	if statusCode == "nil" {
		newRating, anoterStatusCode := updateRating.UpdateUser(*rating, ratingBody, &mongoNewConnection)

		if anoterStatusCode == "nil" {
			ginContext.IndentedJSON(http.StatusOK, newRating)
			return
		} else {
			ginContext.IndentedJSON(http.StatusInternalServerError, newRating)
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
	// defer mongoNewConnection.Disconnect()
	ginContext.IndentedJSON(http.StatusOK, res)
}
