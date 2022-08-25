package routes

import (
	"github.com/gin-goinic.gin"
	controller "golang-restaurant-management/controllers"
)
func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Get("/orders",controller.GetOrders())
	incomingRoutes.Get("/orders/:order_id",controller.GetOrder())
	incomingRoutes.Post("/orders",controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id",controller.UpdateOrder()) 
}