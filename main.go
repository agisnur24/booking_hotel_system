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

	floorRepository := repository.NewFloorRepository()
	floorService := service.NewFloorService(floorRepository, db, validate)
	floorController := controller.NewFloorController(floorService)

	guestRepository := repository.NewGuestRepository()
	guestService := service.NewGuestService(guestRepository, db, validate)
	guestController := controller.NewGuestController(guestService)

	discountRepository := repository.NewDiscountRepository()
	discountService := service.NewDiscountService(discountRepository, db, validate)
	discountController := controller.NewDiscountController(discountService)

	facilityRepository := repository.NewFacilityRepository()
	facilityService := service.NewFacilityService(facilityRepository, db, validate)
	facilityController := controller.NewFacilityController(facilityService)

	meetingRoomRepository := repository.NewMeetingRoomRepository()
	meetingRoomService := service.NewMeetingRoomService(meetingRoomRepository, db, validate)
	meetingRoomController := controller.NewMeetingRoomController(meetingRoomService)

	router := routers.NewUserRouter(userController)
	router = routers.NewRoleRouter(roleController)
	router = routers.NewHotelRouter(hotelController)
	router = routers.NewGuestRouter(guestController)
	router = routers.NewDiscountRouter(discountController)
	router = routers.NewFloorRouter(floorController)
	router = routers.NewFacilityRouter(facilityController)
	router = routers.NewMeetingRoomRouter(meetingRoomController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
