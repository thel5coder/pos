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

type ProductHandler struct {
	Handler
}

func NewProductHandler(handler Handler) handlers.IProductHandler {
	return &ProductHandler{Handler: handler}
}

func (h ProductHandler) Browse(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	uc := usecase.NewProductUseCase(h.UcContract)
	res, pagination, err := uc.Browse(search, orderBy, sort, page, limit)

	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

func (h ProductHandler) ReadByID(ctx *fiber.Ctx) (err error) {
	ID := ctx.Params("id")

	uc := usecase.NewProductUseCase(h.UcContract)
	res, err := uc.ReadBy("p.id", "=", ID)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err)).Send(ctx)
}

func (h ProductHandler) Add(ctx *fiber.Ctx) (err error) {
	req := new(requests.ProductRequest)
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
	uc := usecase.NewProductUseCase(h.UcContract)
	res, err := uc.Add(req)
	if err != nil {
		h.UcContract.TX.Rollback()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UcContract.TX.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(res, err)).Send(ctx)
}

func (h ProductHandler) Edit(ctx *fiber.Ctx) (err error) {
	ID := ctx.Params("id")
	req := new(requests.ProductRequest)
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
	uc := usecase.NewProductUseCase(h.UcContract)
	res, err := uc.Edit(req, ID)
	if err != nil {
		h.UcContract.TX.Rollback()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UcContract.TX.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(res, err)).Send(ctx)
}

func (h ProductHandler) Delete(ctx *fiber.Ctx) (err error) {
	ID := ctx.Params("id")

	// Database processing
	h.UcContract.TX, err = h.DB.Begin()
	if err != nil {
		h.UcContract.TX.Rollback()
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	uc := usecase.NewProductUseCase(h.UcContract)
	err = uc.Delete(ID)
	if err != nil {
		h.UcContract.TX.Rollback()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UcContract.TX.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(nil, err)).Send(ctx)
}
