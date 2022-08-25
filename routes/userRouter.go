package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controllers"

)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Get("/users",controller.GetUsers())
	incomingRoutes.Get("/users/:user_id",controller.GetUser())
	incomingRoutes.Post("/users/signup",controller.SignUp())
	incomingRoutes.Post("/users/login",controller.Login())
}