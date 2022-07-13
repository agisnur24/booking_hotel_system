package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewFloorRouter(floor controller.FloorController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/floors", floor.FindAll)
	router.GET("/api/floors/:floorId", floor.FindById)
	router.POST("/api/floors", floor.Create)
	router.PUT("/api/floors/:floorId", floor.Update)
	router.DELETE("/api/floors/:floorId", floor.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
