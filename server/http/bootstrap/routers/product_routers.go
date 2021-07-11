package routers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/server/http/handlers"
	"majoo-test/server/http/middlewares"
)

type ProductRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func NewProductRouters(routeGroup fiber.Router, handler handlers.Handler) *ProductRouters {
	return &ProductRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (router *ProductRouters) RegisterRouter() {
	handler := handlers.NewProductHandler(router.Handler)
	jwtMiddleware := middlewares.NewJwtMiddleware(router.Handler.UcContract)

	productRoutes := router.RouteGroup.Group("/product")
	productRoutes.Use(jwtMiddleware.Use)
	productRoutes.Get("", handler.Browse)
	productRoutes.Get("/:id", handler.ReadByID)
	productRoutes.Put("/:id", handler.Edit)
	productRoutes.Post("", handler.Add)
	productRoutes.Delete("/:id", handler.Delete)
}
