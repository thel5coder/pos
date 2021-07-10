package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"majoo-test/domain/handlers"
	"majoo-test/domain/requests"
	"majoo-test/pkg/response"
	"majoo-test/usecase"
	"strconv"
)

type UserHandler struct {
	Handler
}

func NewUserHandler(handler Handler) handlers.IUserHandler {
	return &UserHandler{Handler: handler}
}

func (h UserHandler) Browse(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	uc := usecase.NewUserUseCase(h.UcContract)
	res, pagination, err := uc.Browse(search, orderBy, sort, limit, page)

	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

func (h UserHandler) ReadByID(ctx *fiber.Ctx) (err error) {
	ID := ctx.Params("id")
	uc := usecase.NewUserUseCase(h.UcContract)
	res, err := uc.ReadBy("u.id", "=", ID)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err)).Send(ctx)
}

func (h UserHandler) Add(ctx *fiber.Ctx) (err error) {
	req := new(requests.UserRequest)
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
	uc := usecase.NewUserUseCase(h.UcContract)
	res, err := uc.Add(req)
	if err != nil {
		h.UcContract.TX.Rollback()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UcContract.TX.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(res, err)).Send(ctx)
}

func (h UserHandler) Edit(ctx *fiber.Ctx) (err error) {
	ID := ctx.Params("id")
	req := new(requests.UserRequest)
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
	uc := usecase.NewUserUseCase(h.UcContract)
	res, err := uc.Edit(req, ID)
	if err != nil {
		h.UcContract.TX.Rollback()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UcContract.TX.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(res, err)).Send(ctx)
}

func (h UserHandler) Delete(ctx *fiber.Ctx) (err error) {
	ID := ctx.Params("id")

	// Database processing
	h.UcContract.TX, err = h.DB.Begin()
	if err != nil {
		h.UcContract.TX.Rollback()
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	uc := usecase.NewUserUseCase(h.UcContract)
	err = uc.Delete(ID)
	if err != nil {
		h.UcContract.TX.Rollback()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UcContract.TX.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(nil, err)).Send(ctx)
}
