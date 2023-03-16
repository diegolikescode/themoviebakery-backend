package getUser

import (
	"context"
	"log"
	"themoviebakery/config"
	models "themoviebakery/models"
)

func GetUserByEmail(email string, dbConnection *config.PostgresConn) (*models.UserEssentialData, string) {
	statusCode := make(chan string, 1)

	ctx := context.Background()
	query := "SELECT * FROM Users WHERE email = '" + email + "'"
	rows, err := dbConnection.DbConn.QueryContext(ctx, query)

	var user models.UserEssentialData
	var users []models.UserEssentialData
	for rows.Next() {
		var foundUser models.UserEssentialData
		err = rows.Scan(&foundUser)
		if err != nil {
			break
		}
		users = append(users, foundUser)
	}
	if err != nil || len(users) != 0 {
		log.Println("error finding user")
		statusCode <- "USER_NOT_FOUND_BY_EMAIL_404"
	} else {
		statusCode <- "nil"
	}
	// defer dbConnection.Disconnect()
	return &user, <-statusCode
}
