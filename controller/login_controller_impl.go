package controller

/*import (
	"context"
	"net/http"
)

type LoginControllerImpl struct {
	LoginService loginservice.LoginService
}

func NewLoginController(service loginservice.LoginService) LoginController  {
	return &LoginControllerImpl{
		LoginService: service,
	}
}

func (controller *LoginControllerImpl)Login(ctx context.Context)error  {
	bodyrequest := new(request.LoginRequest)

	if err := ctx.Bind(bodyRequest); err != nil{
		return ctx.JSON(http.StatusBadRequest, &http.response.GeneralResponse{
			Code : http.StatusBadRequest,
			Message : err.Error(),
				Data:    nil,
		})
	}

	if err := controller.LoginService.Login(bodyRequest); err != nil {
		return ctx.JSON(http.StatusUnauthorized, &response.GeneralReponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		})
	}

	tokens, err := authentication.GenerateTokenPair()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.GeneralReponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &response.GeneralReponse{
		Code:    http.StatusOK,
		Message: "access granted",
		Data:    tokens,
	})
}*/
