package handlers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/domain/handlers"
	"majoo-test/pkg/response"
	"majoo-test/usecase"
)

type RoleHandler struct{
	Handler
}

func NewRoleHandler(handler Handler) handlers.IRoleHandler{
	return &RoleHandler{Handler:handler}
}

func (h RoleHandler) Browse(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")

	uc := usecase.NewRoleUseCase(h.UcContract)
	res,err := uc.Browse(search)

	return response.NewResponse(response.NewResponseWithOutMeta(res,err)).Send(ctx)
}

