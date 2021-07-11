package handlers

import "github.com/gofiber/fiber/v2"

type IOutletHandler interface {
	Browse(ctx *fiber.Ctx) (err error)
}
