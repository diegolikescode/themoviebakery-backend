package createUser

import (
	"fmt"
	"log"
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
	connectPostgres := config.ConnectPostgres()

	var userBody models.UserTypeInsert
	ginContext.ShouldBindJSON(&userBody)

	_, statusCode := getUser.GetUserByEmail(userBody.Email, &connectPostgres)
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

	query := fmt.Sprint("INSERT INTO Users (userid, email, displayname) VALUES ('", userBody.UserId, "', '", userBody.Email, "', '", userBody.DisplayName, "')")
	_, anotherErr := connectPostgres.DbConn.Exec(query)
	if anotherErr != nil {
		log.Fatal("error creating user")
	}

	defer connectPostgres.Disconnect()

	ginContext.IndentedJSON(http.StatusCreated, bson.M{"newUser": userBody})
}
