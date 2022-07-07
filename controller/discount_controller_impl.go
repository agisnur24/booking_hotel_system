package controller

import (
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
	"github.com/agisnur24/booking_hotel_system.git/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type DiscountControllerImpl struct {
	DiscountService service.DiscountService
}

func NewDiscountController(discountService service.DiscountService) DiscountController {
	return &DiscountControllerImpl{
		DiscountService: discountService,
	}
}

func (controller *DiscountControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	discountCreateRequest := web.DiscountCreateRequest{}
	helper.ReadFromRequestBody(request, &discountCreateRequest)

	discountResponse := controller.DiscountService.Create(request.Context(), discountCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DiscountControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	discountUpdateRequest := web.DiscountUpdateRequest{}
	helper.ReadFromRequestBody(request, &discountUpdateRequest)

	guestId := params.ByName("guestId")
	id, err := strconv.Atoi(guestId)
	helper.PanicIfError(err)

	discountUpdateRequest.Id = id

	discountResponse := controller.DiscountService.Update(request.Context(), discountUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DiscountControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	discountId := params.ByName("discountId")
	id, err := strconv.Atoi(discountId)
	helper.PanicIfError(err)

	controller.DiscountService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DiscountControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	discountId := params.ByName("discountId")
	id, err := strconv.Atoi(discountId)
	helper.PanicIfError(err)

	discountResponse := controller.DiscountService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DiscountControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	discountResponses := controller.DiscountService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
