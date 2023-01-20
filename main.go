package main

import (
	"log"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"themoviebakery/routes"
)

func main() {
	router := SetupRouter()
	log.Fatal(router.Run("localhost:8080"))
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())

	routes.InitUserRoutes(router)
	routes.InitRatingRoutes(router)

	return router
}
