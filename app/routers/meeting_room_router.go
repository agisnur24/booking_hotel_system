package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewMeetingRoomRouter(meetingRoom controller.MeetingRoomController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/meeting_rooms", meetingRoom.FindAll)
	router.GET("/api/meeting_rooms/:meeting_roomId", meetingRoom.FindById)
	router.POST("/api/meeting_rooms", meetingRoom.Create)
	router.PUT("/api/meeting_rooms/:meeting_roomId", meetingRoom.Update)
	router.DELETE("/api/meeting_rooms/:meeting_roomId", meetingRoom.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
