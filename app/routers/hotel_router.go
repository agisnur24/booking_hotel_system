package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewHotelRouter(hotel controller.HotelController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/hotels", hotel.FindAll)
	router.GET("/api/hotels/:hotelId", hotel.FindById)
	router.POST("/api/hotels", hotel.Create)
	router.PUT("/api/hotels/:hotelId", hotel.Update)
	router.DELETE("/api/hotels/:hotelId", hotel.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
