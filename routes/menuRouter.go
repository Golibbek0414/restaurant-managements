package routes

import (
	"github.com/gin-goinic.gin"
	controller "golang-restaurant-management/controllers"
)
func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Get("/menus",controller.GetMenus())
	incomingRoutes.Get("/menus/:menu_id",controller.GetMenu())
	incomingRoutes.Post("/menus",controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id",controller.UpdateMenu()) 
}