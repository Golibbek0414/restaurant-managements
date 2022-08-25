package routes

import (
	"github.com/gin-gonic.gin"
	controller "golang-restaurant-management/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Get("/foods",controller.GetFoods())
	incomingRoutes.Get("/foods/:food_id",controller.GetFood())
	incomingRoutes.Post("/foods",controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id",controller.UpdateFood)

}

