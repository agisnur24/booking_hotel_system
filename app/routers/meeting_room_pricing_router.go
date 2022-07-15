package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewMeetingRoomPricingRouter(meetingRoomPricing controller.MeetingRoomPricingController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/meeting_room_pricings", meetingRoomPricing.FindAll)
	router.GET("/api/meeting_room_pricings/:meeting_room_pricingId", meetingRoomPricing.FindById)
	router.POST("/api/meeting_room_pricings", meetingRoomPricing.Create)
	router.PUT("/api/meeting_room_pricings/:meeting_room_pricingId", meetingRoomPricing.Update)
	router.DELETE("/api/meeting_room_pricings/:meeting_room_pricingId", meetingRoomPricing.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
