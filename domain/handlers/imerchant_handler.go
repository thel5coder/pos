package handlers

import "github.com/gofiber/fiber/v2"

type IMerchantHandler interface {
	Browse(ctx *fiber.Ctx) (err error)
}
