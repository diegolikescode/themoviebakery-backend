package routes

import (
	createRating "themoviebakery/controllers/rating-controllers/create"
	getRatingHandler "themoviebakery/handlers/rating-handlers/get"

	"github.com/gin-gonic/gin"
)

func InitRatingRoutes(router *gin.Engine) {
	groupRoute := router.Group("/api/v1")
	groupRoute.POST("/rating", createRating.CreateRating)
	groupRoute.GET("/rating", getRatingHandler.GetRatingHandler)
}
