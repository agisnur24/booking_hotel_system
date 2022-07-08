package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewFacilityRouter(facility controller.FacilityController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/facilities", facility.FindAll)
	router.GET("/api/facilities/:facilityId", facility.FindById)
	router.POST("/api/facilities", facility.Create)
	router.PUT("/api/facilities/:facilityId", facility.Update)
	router.DELETE("/api/facilities/:facilityId", facility.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
