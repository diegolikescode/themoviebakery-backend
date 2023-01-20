package routes

import (
	createRating "themoviebakery/controllers/rating-controllers/create"

	"github.com/gin-gonic/gin"
)

func InitRatingRoutes(router *gin.Engine) {
	groupRoute := router.Group("/api/v1")
	groupRoute.POST("/rating", createRating.CreateRating)
}
