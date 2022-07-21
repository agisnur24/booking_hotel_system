package controller

import (
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
	"github.com/agisnur24/booking_hotel_system.git/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func NewRoleController(roleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

func (controller *RoleControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleCreateRequest := web.RoleCreateRequest{}
	helper.ReadFromRequestBody(request, &roleCreateRequest)

	roleResponse := controller.RoleService.Create(request.Context(), roleCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleUpdateRequest := web.RoleUpdateRequest{}
	helper.ReadFromRequestBody(request, &roleUpdateRequest)

	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)

	roleUpdateRequest.Id = id

	roleResponse := controller.RoleService.Update(request.Context(), roleUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)

	controller.RoleService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("roleId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	customerResponse := controller.RoleService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleResponses := controller.RoleService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
