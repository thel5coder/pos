package handlers

import (
	"database/sql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"majoo-test/pkg/jwe"
	"majoo-test/pkg/jwt"
	"majoo-test/usecase"
)

type Handler struct {
	App           *fiber.App
	UcContract    *usecase.Contract
	DB            *sql.DB
	Validate      *validator.Validate
	Translator    ut.Translator
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
}
