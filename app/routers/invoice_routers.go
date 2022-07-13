package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewInvoiceRouter(invoice controller.InvoiceController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/invoices", invoice.FindAll)
	router.GET("/api/invoices/:invoiceId", invoice.FindById)
	router.POST("/api/invoices", invoice.Create)
	router.PUT("/api/invoices/:invoiceId", invoice.Update)
	router.DELETE("/api/invoices/:invoiceId", invoice.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
