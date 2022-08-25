package routes

import (
	"github.com/gin-goinic.gin"
	controller "golang-restaurant-management/controllers"
)
func OrderItemsRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Get("/orderItems",controller.GetOrderItems())
	incomingRoutes.Get("/orderItems/:orderItem_id",controller.GetOrderItem())
	incomingRoutes.Get("/orderItmes-order/:order_id",controller.GetOrderItemsByOrder())
	incomingRoutes.Post("/orderItems",controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id",controller.UpdateOrderItem()) 
}