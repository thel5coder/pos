package handlers

import "github.com/gofiber/fiber/v2"

type IProductHandler interface {
	Browse(ctx *fiber.Ctx) (err error)

	ReadByID(ctx *fiber.Ctx) (err error)

	Add(ctx *fiber.Ctx) (err error)

	Edit(ctx *fiber.Ctx) (err error)

	Delete(ctx *fiber.Ctx) (err error)
}
