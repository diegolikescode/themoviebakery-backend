package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"themoviebakery/config"
	createUser "themoviebakery/controllers/user-controllers/create"
)

type user struct {
	UserID          string `json:"userId"`
	Email           string `json:"email"`
	DisplayName     string `json:"displaName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func getUsers(c *gin.Context) {
	collection := config.ConnectMongo()
	createUser.CreateUser(collection)

	c.IndentedJSON(http.StatusOK, collection)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)

	router.Run("localhost:3030")
}
