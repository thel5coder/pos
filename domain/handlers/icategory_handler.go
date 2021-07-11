package handlers

import "github.com/gofiber/fiber/v2"

type ICategoryHandler interface {
	Browse(ctx *fiber.Ctx) (err error)
}
