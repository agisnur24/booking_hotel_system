package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewDiscountRouter(discount controller.DiscountController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/discounts", discount.FindAll)
	router.GET("/api/discounts/:discountId", discount.FindById)
	router.POST("/api/discounts", discount.Create)
	router.PUT("/api/discounts/:discountId", discount.Update)
	router.DELETE("/api/discounts/:discountId", discount.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
