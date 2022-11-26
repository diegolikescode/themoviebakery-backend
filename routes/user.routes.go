package routes

import (
	createUser "themoviebakery/controllers/user-controllers/create"
	getUser "themoviebakery/controllers/user-controllers/get"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.Engine) {
	groupRoute := router.Group("/api/v1") // .Use(middleware.Auth())
	groupRoute.POST("/user", createUser.CreateUser)
	groupRoute.GET("/user-by-email", getUser.GetUserByEmail)
	groupRoute.GET("/user-by-id", getUser.GetUserById)
	// groupRoute.PUT("/user", createUser.UpdateUser)
	// groupRoute.DELETE("/user", createUser.DeleteUser)
}
