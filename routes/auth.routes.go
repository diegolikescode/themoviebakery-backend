package routes

import (
	loginHandler "themoviebakery/handlers/auth-handlers/login"
	// updateUser "themoviebakery/controllers/auth-controllers"
	// deleteUserHandler "themoviebakery/handlers/auth-handlers"
	// getUserHandler "themoviebakery/handlers/auth-handlers"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(router *gin.Engine) {
	groupRoute := router.Group("/api/v1") // .Use(middleware.Auth())
	groupRoute.POST("/auth", loginHandler.LoginHandler)
}
