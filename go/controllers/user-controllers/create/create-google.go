package createUser

import (
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
	postgresConnection := config.ConnectPostgres()

	var userBody models.UserTypeInsert
	ginContext.ShouldBindJSON(&userBody)

	user, statusCode := getUser.GetUserByEmail(userBody.Email, &postgresConnection)
	fmt.Println(statusCode)
	if statusCode == "nil" {
		ginContext.IndentedJSON(http.StatusOK, user)
		defer postgresConnection.Disconnect()
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

	query := fmt.Sprint("INSERT INTO Users (userid, email, displayname) VALUES (", userBody.UserId, ", ", userBody.Email, ", ", userBody.DisplayName)
	_, anotherErr := postgresConnection.DbConn.Exec(query)
	if anotherErr != nil {
		panic(err)
	}

	defer postgresConnection.Disconnect()

	ginContext.IndentedJSON(http.StatusCreated, bson.M{"newUser": userBody})
}
