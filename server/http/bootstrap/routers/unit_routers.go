package routers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/server/http/handlers"
	"majoo-test/server/http/middlewares"
)

type UnitRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func NewUnitRouters(routeGroup fiber.Router, handler handlers.Handler) *UnitRouters {
	return &UnitRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (router *UnitRouters) RegisterRouter() {
	handler := handlers.NewUnitHandler(router.Handler)
	jwtMiddleware := middlewares.NewJwtMiddleware(router.Handler.UcContract)

	unitRoutes := router.RouteGroup.Group("/unit")
	unitRoutes.Use(jwtMiddleware.Use)
	unitRoutes.Get("", handler.Browse)
}
