package handlers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/domain/handlers"
	"majoo-test/pkg/response"
	"majoo-test/usecase"
)

type CategoryHandler struct{
	Handler
}

func NewCategoryHandler(handler Handler) handlers.ICategoryHandler{
	return &CategoryHandler{Handler:handler}
}

func (h CategoryHandler) Browse(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")

	uc := usecase.NewCategoryUseCase(h.UcContract)
	res,err := uc.Browse(search)

	return response.NewResponse(response.NewResponseWithOutMeta(res,err)).Send(ctx)
}

