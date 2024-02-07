package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/golangcompany/JWT-Authentication/controllers"
	"github.com/golangcompany/JWT-Authentication/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.UserAuthenticate())
	incomingRoutes.GET("/usersdata", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
