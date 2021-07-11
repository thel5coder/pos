package routers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/server/http/handlers"
	"majoo-test/server/http/middlewares"
)

type CategoryRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func NewCategoryRouters(routeGroup fiber.Router, handler handlers.Handler) *CategoryRouters {
	return &CategoryRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (router *CategoryRouters) RegisterRouter() {
	handler := handlers.NewCategoryHandler(router.Handler)
	jwtMiddleware := middlewares.NewJwtMiddleware(router.Handler.UcContract)

	categoryRoutes := router.RouteGroup.Group("/category")
	categoryRoutes.Use(jwtMiddleware.Use)
	categoryRoutes.Get("", handler.Browse)
}
