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

// AuthRoutes sets up the authentication routes for the incomingRoutes.
// It adds the UserAuthenticate middleware to authenticate the user.
// It also registers two GET routes: "/usersdata" and "/users/:user_id".
// The "/usersdata" route is handled by the GetUsers function in the controller package.
// The "/users/:user_id" route is handled by the GetUser function in the controller package.
//
// Parameters:
// - incomingRoutes: A pointer to the gin.Engine instance to which the routes will be added.
func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.UserAuthenticate())
	incomingRoutes.GET("/usersdata", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
