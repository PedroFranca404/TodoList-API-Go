package router

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()

	// Initialize routes
	initRoutes(router)

	router.Run(":8080")
}
