package bootstrap

import (
	"database/sql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"majoo-test/usecase"
)

type Bootstrap struct {
	App        *fiber.App
	DB         *sql.DB
	UcContract usecase.Contract
	Validator  *validator.Validate
	Translator ut.Translator
}
