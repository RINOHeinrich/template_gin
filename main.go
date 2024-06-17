package main

import (
	routes "github.com/RINOHeinrich/template_gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "8000"

	router := gin.New()

	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
