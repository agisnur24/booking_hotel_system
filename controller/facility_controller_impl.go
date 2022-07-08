package controller

import (
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
	"github.com/agisnur24/booking_hotel_system.git/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type FacilityControllerImpl struct {
	FacilityService service.FacilityService
}

func NewFacilityController(facilityService service.FacilityService) FacilityController {
	return &FacilityControllerImpl{
		FacilityService: facilityService,
	}
}

func (controller *FacilityControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	facilityCreateRequest := web.FacilityCreateRequest{}
	helper.ReadFromRequestBody(request, &facilityCreateRequest)

	facilityResponse := controller.FacilityService.Create(request.Context(), facilityCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   facilityResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FacilityControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	facilityUpdateRequest := web.FacilityUpdateRequest{}
	helper.ReadFromRequestBody(request, &facilityUpdateRequest)

	facilityId := params.ByName("facilityId")
	id, err := strconv.Atoi(facilityId)
	helper.PanicIfError(err)

	facilityUpdateRequest.Id = id

	facilityResponse := controller.FacilityService.Update(request.Context(), facilityUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   facilityResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FacilityControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	facilityId := params.ByName("facilityId")
	id, err := strconv.Atoi(facilityId)
	helper.PanicIfError(err)

	controller.FacilityService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FacilityControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	facilityId := params.ByName("facilityId")
	id, err := strconv.Atoi(facilityId)
	helper.PanicIfError(err)

	facilityResponse := controller.FacilityService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   facilityResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FacilityControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	facilityResponses := controller.FacilityService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   facilityResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
