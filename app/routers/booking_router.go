package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewBookingRouter(booking controller.BookingController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/bookings", booking.FindAll)
	router.GET("/api/bookings/:bookingId", booking.FindById)
	router.POST("/api/bookings", booking.Create)
	router.PUT("/api/bookings/:bookingId", booking.Update)
	router.DELETE("/api/booking/:bookingId", booking.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
