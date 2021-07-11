package handlers

import "github.com/gofiber/fiber/v2"

type IRoleHandler interface {
	Browse(ctx *fiber.Ctx) (err error)
}
