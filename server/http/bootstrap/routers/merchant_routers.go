package routers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/server/http/handlers"
	"majoo-test/server/http/middlewares"
)

type MerchantRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func NewMerchantRouters(routeGroup fiber.Router, handler handlers.Handler) *MerchantRouters {
	return &MerchantRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (router MerchantRouters) RegisterRouter() {
	handler := handlers.NewMerchantHandler(router.Handler)
	jwtMiddleware := middlewares.NewJwtMiddleware(router.Handler.UcContract)

	merchantRoutes := router.RouteGroup.Group("/merchant")
	merchantRoutes.Use(jwtMiddleware.Use)
	merchantRoutes.Get("", handler.Browse)
}
