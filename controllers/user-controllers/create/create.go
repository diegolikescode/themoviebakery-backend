package createUser

import (
	"context"
	"net/http"
	config "themoviebakery/config"
	getUser "themoviebakery/controllers/user-controllers/get"
	"time"

	models "themoviebakery/models"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(ginContext *gin.Context) {
	mongoNewConnection := config.ConnectMongo("users")

	var userBody models.UserTypeInsert
	ginContext.ShouldBindJSON(&userBody)

	_, statusCode := getUser.GetUserByEmail(userBody.Email, &mongoNewConnection)
	if statusCode == "nil" {
		ginContext.IndentedJSON(http.StatusConflict, bson.M{"message": "user already exists"})
		return
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

	userBody.UserId = shortId
	userBody.CreatedAt = time.Now()
	userBody.UpdatedAt = time.Now()

	res, err := mongoNewConnection.Collection.InsertOne(context.TODO(), userBody)
	if err != nil {
		panic(err)
	}

	defer mongoNewConnection.Disconnect()

	ginContext.IndentedJSON(http.StatusOK, res)
}
