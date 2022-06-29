package app

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func UserRouter(user controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users", user.FindAll)
	router.GET("/api/users/:userEmail", user.FindByEmail)
	router.POST("/api/users", user.Create)
	router.PUT("/api/users/:userEmail", user.Update)
	router.DELETE("/api/users/:userEmail", user.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
