package createUser

import (
	"context"
	"fmt"
	"net/http"
	config "themoviebakery/config"
	getUser "themoviebakery/controllers/user-controllers/get"
	models "themoviebakery/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateAndLogUserGoogle(ginContext *gin.Context) {
	// UserId      string    `json:"userId"`
	// Email       string    `json:"email" validate:"required"`
	// DisplayName string    `json:"displayName" validate:"required"`
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
	mongoNewConnection := config.ConnectMongo("users")

	var userBody models.UserTypeInsert
	ginContext.ShouldBindJSON(&userBody)

	user, statusCode := getUser.GetUserByEmail(userBody.Email, &mongoNewConnection)
	fmt.Println(statusCode)
	if statusCode == "nil" {
		ginContext.IndentedJSON(http.StatusOK, user)
		defer mongoNewConnection.Disconnect()
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

	insertedId, err := mongoNewConnection.Collection.InsertOne(context.TODO(), userBody)
	if err != nil {
		panic(err)
	}

	defer mongoNewConnection.Disconnect()

	ginContext.IndentedJSON(http.StatusCreated, bson.M{"newUser": userBody, "insertedId": insertedId})
}
