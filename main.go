package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	UserID          string `json:"userId"`
	Email           string `json:"email"`
	DisplayName     string `json:"displaName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func getUsers(c *gin.Context) {
	var users = []user{
		{UserID: "aueauhe", Email: "aseuas", DisplayName: "UAEHUAEHUA", Password: "ahursfdhgw", ConfirmPassword: "HURSDFHUGSDGUH"},
		{UserID: "aueauhe", Email: "aseuas", DisplayName: "UAEHUAEHUA", Password: "ahursfdhgw", ConfirmPassword: "HURSDFHUGSDGUH"},
		{UserID: "aueauhe", Email: "aseuas", DisplayName: "UAEHUAEHUA", Password: "fgf", ConfirmPassword: "jhjjhjh"},
		{UserID: "aueauhe", Email: "aseuas", DisplayName: "UAEHUAEHUA", Password: "ahursfdhgw", ConfirmPassword: "sadasasas"},
	}

	// append(users, )

	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)

	router.Run("localhost:3030")
}
