package usecase

import (
	"majoo-test/domain/requests"
	"majoo-test/domain/view_models"
)

type IAuthenticationUseCase interface {
	Login(req *requests.LoginRequest) (res view_models.LoginVm, err error)

	GenerateJWT(issuer, payload string) (res view_models.LoginVm, err error)
}
