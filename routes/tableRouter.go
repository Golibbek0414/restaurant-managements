package routes

import (
	"github.com/gin-goinic.gin"
	controller "golang-restaurant-management/controllers"
)
func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Get("/tables",controller.GetTables())
	incomingRoutes.Get("/tables/:table_id",controller.GetTable())
	incomingRoutes.Post("/tables",controller.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id",controller.UpdateTable()) 
}