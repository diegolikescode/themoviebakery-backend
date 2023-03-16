package getUser

import (
	"log"
	"themoviebakery/config"
	models "themoviebakery/models"
)

func GetUserById(id string, dbConnection *config.PostgresConn) (*models.UserTypeFullIdPrimitive, string) {
	statusCode := make(chan string, 1)

	var user models.UserTypeFullIdPrimitive
	query := ""
	_, err := dbConnection.DbConn.Exec(query)
	if err != nil {
		log.Println("Trying to parse user's object id (string) to mongo's ObjectId, Invalid ID =>", err)
		statusCode <- "ERROR_FINDING_USER_BY_ID_404"
		return nil, <-statusCode
	} else {
		statusCode <- "nil"
	}

	return &user, <-statusCode
}
