package handlers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/domain/handlers"
	"majoo-test/pkg/response"
	"majoo-test/usecase"
)

type OutletHandler struct{
	Handler
}

func NewOutletHandler(handler Handler) handlers.IOutletHandler{
	return &OutletHandler{Handler:handler}
}

func (h OutletHandler) Browse(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")

	uc := usecase.NewOutletUseCase(h.UcContract)
	res,err := uc.Browse(search)

	return response.NewResponse(response.NewResponseWithOutMeta(res,err)).Send(ctx)
}

