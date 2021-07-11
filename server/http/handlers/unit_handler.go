package handlers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/domain/handlers"
	"majoo-test/pkg/response"
	"majoo-test/usecase"
)

type UnitHandler struct {
	Handler
}

func NewUnitHandler(handler Handler) handlers.IUnitHandler {
	return &UnitHandler{Handler: handler}
}

func (h UnitHandler) Browse(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")

	uc := usecase.NewUnitUseCase(h.UcContract)
	res,err := uc.Browse(search)

	return response.NewResponse(response.NewResponseWithOutMeta(res,err)).Send(ctx)
}
