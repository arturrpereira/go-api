package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	router.Run()
}
