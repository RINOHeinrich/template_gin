package main

import (
	routes "github.com/RINOBE/gestion_de_livre/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "8000"
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	/*router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})
	*/
	router.Run(":" + port)
}
