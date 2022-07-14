package main

import (
	"github.com/agisnur24/booking_hotel_system.git/app"
	"github.com/agisnur24/booking_hotel_system.git/app/routers"
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/middleware"
	"github.com/agisnur24/booking_hotel_system.git/repository"
	"github.com/agisnur24/booking_hotel_system.git/service"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	hotelRepository := repository.NewHotelRepository()
	hotelService := service.NewHotelService(hotelRepository, db, validate)
	hotelController := controller.NewHotelController(hotelService)

	invoiceRepository := repository.NewInvoiceRepository()
	invoiceService := service.NewInvoiceService(invoiceRepository, db, validate)
	invoiceController := controller.NewInvoiceController(invoiceService)

	bookingRepository := repository.NewBookingRepository()
	bookingService := service.NewBookingService(bookingRepository, db, validate)
	bookingController := controller.NewBookingController(bookingService)

	discountRepository := repository.NewDiscountRepository()
	discountService := service.NewDiscountService(discountRepository, db, validate)
	discountController := controller.NewDiscountController(discountService)

	router := routers.NewUserRouter(userController)
	router = routers.NewRoleRouter(roleController)
	router = routers.NewHotelRouter(hotelController)
	router = routers.NewInvoiceRouter(invoiceController)
	router = routers.NewBookingRouter(bookingController)
	router = routers.NewDiscountRouter(discountController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
