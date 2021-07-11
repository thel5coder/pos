package middlewares

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"majoo-test/pkg/functioncaller"
	jwtPkg "majoo-test/pkg/jwt"
	"majoo-test/pkg/logruslogger"
	"majoo-test/pkg/messages"
	"majoo-test/pkg/response"
	"majoo-test/pkg/str"
	"majoo-test/usecase"
	"strings"
	"time"
)

type JwtMiddleware struct {
	*usecase.Contract
}

func NewJwtMiddleware(contract *usecase.Contract) JwtMiddleware {
	return JwtMiddleware{Contract: contract}
}

func (m JwtMiddleware) Use(ctx *fiber.Ctx) error {
	claims := &jwtPkg.CustomClaims{}

	//check header is present or not
	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "middleware-jwt-checkHeader")
		return response.NewResponse(response.NewResponseWithOutMeta(nil, errors.New(messages.Unauthorized))).Send(ctx)
	}

	//check claims and signing method
	token := strings.Replace(header, "Bearer ", "", -1)
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			logruslogger.Log(logruslogger.WarnLevel, messages.UnexpectedSigningMethod, functioncaller.PrintFuncName(), "middleware-jwt-checkSigningMethod")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := []byte(m.JwtCredential.TokenSecret)
		return secret, nil
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "middleware-jwt-checkClaims")
		return response.NewResponse(response.NewResponseWithOutMeta(nil, errors.New(messages.Unauthorized))).Send(ctx)
	}

	//check token live time
	if claims.ExpiresAt < time.Now().Unix() {
		logruslogger.Log(logruslogger.WarnLevel, messages.ExpiredToken, functioncaller.PrintFuncName(), "middleware-jwt-checkTokenLiveTime")
		return response.NewResponse(response.NewResponseWithOutMeta(nil, errors.New(messages.Unauthorized))).Send(ctx)
	}

	//jwe roll back encrypted id
	jweRes, err := m.JweCredential.Rollback(claims.Payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-jwe-rollback")
		return response.NewResponse(response.NewResponseWithOutMeta(nil, errors.New(messages.Unauthorized))).Send(ctx)
	}
	if jweRes == nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "pkg-jwe-resultNil")
		return response.NewResponse(response.NewResponseWithOutMeta(nil, errors.New(messages.Unauthorized))).Send(ctx)
	}

	//set id to uce case contract
	roleID := fmt.Sprintf("%v", jweRes["roleID"])
	m.Contract.UserID = fmt.Sprintf("%v", jweRes["id"])
	m.Contract.RoleID = str.StringToInt(roleID)
	m.Contract.MerchantID = fmt.Sprintf("%v", jweRes["merchantID"])

	var userLoggedIn map[string]interface{}
	err = m.Redis.GetFromRedis(token, &userLoggedIn)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "pkg-redis-getFromRedis")
		return response.NewResponse(response.NewResponseWithOutMeta(nil, errors.New(messages.Unauthorized))).Send(ctx)
	}

	return ctx.Next()
}
