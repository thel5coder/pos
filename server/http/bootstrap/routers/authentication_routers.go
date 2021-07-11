package routers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/server/http/handlers"
)

type AuthenticationRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func NewAuthenticationRouters(routeGroup fiber.Router,handler handlers.Handler) *AuthenticationRouters{
	return &AuthenticationRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func(router *AuthenticationRouters) RegisterRouter(){
	handler := handlers.NewAuthenticationHandler(router.Handler)

	authenticationRoutes := router.RouteGroup.Group("/auth")
	authenticationRoutes.Post("/login",handler.Login)
}
