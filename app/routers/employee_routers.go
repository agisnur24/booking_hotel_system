package routers

import (
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewEmployeeRouter(employee controller.EmployeeController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/employees", employee.FindAll)
	router.GET("/api/employees/:employeeId", employee.FindById)
	router.POST("/api/employees", employee.Create)
	router.PUT("/api/employees/:employeeId", employee.Update)
	router.DELETE("/api/employees/:employeeId", employee.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
