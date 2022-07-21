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
	"log"
	"net/http"
	"os"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	bookingRepository := repository.NewBookingRepository()
	bookingService := service.NewBookingService(bookingRepository, db, validate)
	bookingController := controller.NewBookingController(bookingService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	hotelRepository := repository.NewHotelRepository()
	hotelService := service.NewHotelService(hotelRepository, db, validate)
	hotelController := controller.NewHotelController(hotelService)

	discountRepository := repository.NewDiscountRepository()
	discountService := service.NewDiscountService(discountRepository, db, validate)
	discountController := controller.NewDiscountController(discountService)

	meetingRoomRepository := repository.NewMeetingRoomRepository()
	meetingRoomService := service.NewMeetingRoomService(meetingRoomRepository, db, validate)
	meetingRoomController := controller.NewMeetingRoomController(meetingRoomService)

	invoiceRepository := repository.NewInvoiceRepository()
	invoiceService := service.NewInvoiceService(invoiceRepository, db, validate)
	invoiceController := controller.NewInvoiceController(invoiceService)

	router := routers.NewBookingRouter(bookingController)
	router = routers.NewHotelRouter(hotelController)
	router = routers.NewRoleRouter(roleController)
	router = routers.NewUserRouter(userController)
	router = routers.NewDiscountRouter(discountController)
	router = routers.NewMeetingRoomRouter(meetingRoomController)
	router = routers.NewInvoiceRouter(invoiceController)

	server := http.Server{
		//Addr:    "localhost:3000",
		Addr:    "https://booking-management-system-fp.herokuapp.com/" + os.Getenv("PORT") + "/",
		Handler: middleware.NewAuthMiddleware(router),
	}
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
