package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewUserRouter(user controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users", user.FindAll)
	router.GET("/api/users/:userId", user.FindById)
	router.POST("/api/users", user.Create)
	router.PUT("/api/users/:userId", user.Update)
	router.DELETE("/api/users/:userId", user.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
