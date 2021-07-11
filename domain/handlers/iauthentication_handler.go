package handlers

import "github.com/gofiber/fiber/v2"

type IAuthenticationHandler interface {
	Login(ctx *fiber.Ctx) (err error)
}
