package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/golangcompany/JWT-Authentication/routes"
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
