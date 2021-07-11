package handlers

import "github.com/gofiber/fiber/v2"

type IUnitHandler interface {
	Browse(ctx *fiber.Ctx)(err error)
}
