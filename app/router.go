package app

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(user controller.UserController, role controller.RoleController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users", user.FindAll)
	router.GET("/api/users/:userEmail", user.FindByEmail)
	router.POST("/api/users", user.Create)
	router.PUT("/api/users/:userEmail", user.Update)
	router.DELETE("/api/users/:userEmail", user.Delete)

	router.GET("/api/roles", role.FindAll)
	router.GET("/api/roles/:roleName", role.FindByRoleName)
	router.POST("/api/roles", role.Create)
	router.PUT("/api/roles/:roleName", role.Update)
	router.DELETE("/api/roles/:roleName", role.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
