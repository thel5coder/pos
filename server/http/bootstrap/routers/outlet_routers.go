package routers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/server/http/handlers"
	"majoo-test/server/http/middlewares"
)

type OutletRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func NewOutletRouters(routeGroup fiber.Router,handler handlers.Handler) *OutletRouters{
	return &OutletRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func(router *OutletRouters) RegisterRoute(){
	handler := handlers.NewOutletHandler(router.Handler)
	jwtMiddleware := middlewares.NewJwtMiddleware(router.Handler.UcContract)

	outletRoutes := router.RouteGroup.Group("/outlet")
	outletRoutes.Use(jwtMiddleware.Use)
	outletRoutes.Get("",handler.Browse)
}
