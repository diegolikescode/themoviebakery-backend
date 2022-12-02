package routes

import (
	createUser "themoviebakery/controllers/user-controllers/create"
	updateUser "themoviebakery/controllers/user-controllers/update"
	deleteUserHandler "themoviebakery/handlers/user-handlers/delete"
	getUserHandler "themoviebakery/handlers/user-handlers/get"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.Engine) {
	groupRoute := router.Group("/api/v1") // .Use(middleware.Auth())
	groupRoute.POST("/user", createUser.CreateUser)
	groupRoute.GET("/user-by-email", getUserHandler.HandlerGetUserByEmail)
	groupRoute.GET("/user-by-id", getUserHandler.HandlerGetUserById)
	groupRoute.PUT("/user", updateUser.UpdateUser)
	groupRoute.DELETE("/user", deleteUserHandler.DeleteUserById)
}
