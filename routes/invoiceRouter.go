package routes

import (
	"github.com/gin-goinic.gin"
	controller "golang-restaurant-management/controllers"
)
func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Get("/invoices",controller.GetInvoices())
	incomingRoutes.Get("/invoices/:invoices_id",controller.GetInvoice())
	incomingRoutes.Post("/invoices",controller.CreateInvoice())
	incomingRoutes.PATCH("/inoices/:inoices_id",controller.UpdateInvoice)
}