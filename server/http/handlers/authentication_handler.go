package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"majoo-test/domain/handlers"
	"majoo-test/domain/requests"
	"majoo-test/pkg/response"
	"majoo-test/usecase"
)

type AuthenticationHandler struct{
	Handler
}

func NewAuthenticationHandler(handler Handler) handlers.IAuthenticationHandler{
	return &AuthenticationHandler{Handler:handler}
}

func (h AuthenticationHandler) Login(ctx *fiber.Ctx) (err error) {
	req := new(requests.LoginRequest)
	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.Validate.Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.Translator)).Send(ctx)
	}

	// Database processing
	h.UcContract.TX, err = h.DB.Begin()
	if err != nil {
		h.UcContract.TX.Rollback()
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	uc := usecase.NewAuthenticationUseCase(h.UcContract)
	res, err := uc.Login(req)
	if err != nil {
		h.UcContract.TX.Rollback()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UcContract.TX.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(res, err)).Send(ctx)
}

