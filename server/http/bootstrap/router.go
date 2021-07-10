package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"majoo-test/server/http/bootstrap/routers"
	"majoo-test/server/http/handlers"
)

func (boot Bootstrap) RegisterRoute() {
	handlerType := handlers.Handler{
		App:        boot.App,
		UcContract: &boot.UcContract,
		DB:         boot.DB,
		Validate:   boot.Validator,
		Translator: boot.Translator,
	}

	//route for check health
	rootParentGroup := boot.App.Group("/")
	rootParentGroup.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("it's working")
	})

	userRoutes := routers.NewUserRouters(rootParentGroup, handlerType)
	userRoutes.RegisterRouter()
}
