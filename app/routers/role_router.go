package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRoleRouter(role controller.RoleController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/roles", role.FindAll)
	router.GET("/api/roles/:roleId", role.FindById)
	router.POST("/api/roles", role.Create)
	router.PUT("/api/roles/:roleId", role.Update)
	router.DELETE("/api/roles/:roleId", role.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
