package handlers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/domain/handlers"
	"majoo-test/pkg/response"
	"majoo-test/usecase"
)

type MerchantHandler struct{
	Handler
}

func NewMerchantHandler(handler Handler) handlers.IMerchantHandler{
	return MerchantHandler{Handler:handler}
}

func (h MerchantHandler) Browse(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")

	uc := usecase.NewMerchantUseCase(h.UcContract)
	res,err := uc.Browse(search)

	return response.NewResponse(response.NewResponseWithOutMeta(res,err)).Send(ctx)
}

