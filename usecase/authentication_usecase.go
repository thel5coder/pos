package usecase

import (
	"errors"
	"fmt"
	"majoo-test/domain/requests"
	"majoo-test/domain/usecase"
	"majoo-test/domain/view_models"
	"majoo-test/pkg/functioncaller"
	"majoo-test/pkg/hashing"
	"majoo-test/pkg/logruslogger"
	"majoo-test/pkg/messages"
	"majoo-test/repository/query"
	"os"
)

type AuthenticationUseCase struct {
	*Contract
}

func NewAuthenticationUseCase(contract *Contract) usecase.IAuthenticationUseCase {
	return &AuthenticationUseCase{Contract: contract}
}

func (uc AuthenticationUseCase) Login(req *requests.LoginRequest) (res view_models.LoginVm, err error) {
	q := query.NewUserQuery(uc.DB)
	user, err := q.ReadBy("u.email", "=", req.Email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-readByEmail")
		return res, errors.New(messages.Unauthorized)
	}

	if isPasswordValid := hashing.CheckHashString(req.Password, user.Password()); !isPasswordValid {
		logruslogger.Log(logruslogger.WarnLevel, messages.CredentialDoNotMatch, functioncaller.PrintFuncName(), "hash-check")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	fmt.Println(user.MerchantID())
	//generate jwt payload and encrypted with jwe
	payload := map[string]interface{}{
		"id":         user.Id(),
		"roleID":     user.RoleID(),
		"role":       user.Roles.Name(),
		"merchantID": user.MerchantID(),
	}
	jwePayload, err := uc.JweCredential.GenerateJwePayload(payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-generate-jwe-payload")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	//generate jwt token
	res, err = uc.GenerateJWT(req.Email, jwePayload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-generate-jwt-token")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	userLoggedIn := map[string]interface{}{
		"email": user.Email(),
		"role": map[string]interface{}{
			"id":   user.RoleID(),
			"name": user.Roles.Name(),
		},
	}
	err = uc.Redis.StoreToRedisWithExpired(res.Token, userLoggedIn, os.Getenv("TOKEN_EXP_TIME")+`h`)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-generate-jwt-token")
		return res, err
	}

	return res, nil
}

func (uc AuthenticationUseCase) GenerateJWT(issuer, payload string) (res view_models.LoginVm, err error) {
	res.Token, res.TokenExpiredAt, err = uc.JwtCredential.GetToken(issuer, payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-jwt-getToken")
		return res, err
	}

	res.RefreshToken, res.RefreshTokenExpiredAt, err = uc.JwtCredential.GetRefreshToken(issuer, payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-jwt-getRefreshToken")
		return res, err
	}

	return res, nil
}
