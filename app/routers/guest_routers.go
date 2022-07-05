package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewGuestRouter(guest controller.GuestController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/guests", guest.FindAll)
	router.GET("/api/guests/:guestId", guest.FindById)
	router.POST("/api/guests", guest.Create)
	router.PUT("/api/guests/:guestId", guest.Update)
	router.DELETE("/api/guests/:guestId", guest.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
