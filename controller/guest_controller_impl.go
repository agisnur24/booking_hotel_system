package controller

import (
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
	"github.com/agisnur24/booking_hotel_system.git/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type GuestControllerImpl struct {
	GuestService service.GuestService
}

func NewGuestController(guestService service.GuestService) GuestController {
	return &GuestControllerImpl{
		GuestService: guestService,
	}
}

func (controller *GuestControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	guestCreateRequest := web.GuestCreateRequest{}
	helper.ReadFromRequestBody(request, &guestCreateRequest)

	guestResponse := controller.GuestService.Create(request.Context(), guestCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   guestResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *GuestControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	guestUpdateRequest := web.GuestUpdateRequest{}
	helper.ReadFromRequestBody(request, &guestUpdateRequest)

	guestId := params.ByName("guestId")
	id, err := strconv.Atoi(guestId)
	helper.PanicIfError(err)

	guestUpdateRequest.Id = id

	guestResponse := controller.GuestService.Update(request.Context(), guestUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   guestResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *GuestControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	guestId := params.ByName("guestId")
	id, err := strconv.Atoi(guestId)
	helper.PanicIfError(err)

	controller.GuestService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *GuestControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	guestId := params.ByName("guestId")
	id, err := strconv.Atoi(guestId)
	helper.PanicIfError(err)

	guestResponse := controller.GuestService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   guestResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *GuestControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	guestResponses := controller.GuestService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   guestResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
