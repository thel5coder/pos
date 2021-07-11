package routers

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/server/http/handlers"
	"majoo-test/server/http/middlewares"
)

type UserRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func NewUserRouters(routeGroup fiber.Router, handler handlers.Handler) *UserRouters {
	return &UserRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (router *UserRouters) RegisterRouter() {
	handler := handlers.NewUserHandler(router.Handler)
	jwtMiddleware := middlewares.NewJwtMiddleware(router.Handler.UcContract)

	userRoutes := router.RouteGroup.Group("/user")
	userRoutes.Use(jwtMiddleware.Use)
	userRoutes.Get("", handler.Browse)
	userRoutes.Get("/:id", handler.ReadByID)
	userRoutes.Put("/:id", handler.Edit)
	userRoutes.Post("", handler.Add)
	userRoutes.Delete("/:id", handler.Delete)
}
