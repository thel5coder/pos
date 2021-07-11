package routers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/server/http/handlers"
	"majoo-test/server/http/middlewares"
)

type RoleRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func NewRoleRouters(routeGroup fiber.Router, handler handlers.Handler) *RoleRouters {
	return &RoleRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func(router *RoleRouters) RegisterRouter() {
	handler := handlers.NewRoleHandler(router.Handler)
	jwtMiddleware := middlewares.NewJwtMiddleware(router.Handler.UcContract)

	roleRoutes := router.RouteGroup.Group("/role")
	roleRoutes.Use(jwtMiddleware.Use)
	roleRoutes.Get("",handler.Browse)
}
