package routes

import (
	createUser "themoviebakery/controllers/user-controllers/create"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.Engine) {
	groupRoute := router.Group("/api/v1") // .Use(middleware.Auth())
	groupRoute.POST("/user", createUser.CreateUser)
	// groupRoute.GET("/user", createUser.GetUser)
	// groupRoute.PUT("/user", createUser.UpdateUser)
	// groupRoute.DELETE("/user", createUser.DeleteUser)
}
