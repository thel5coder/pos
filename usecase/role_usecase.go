package usecase

import (
	"majoo-test/domain/usecase"
	"majoo-test/domain/view_models"
	"majoo-test/pkg/functioncaller"
	"majoo-test/pkg/logruslogger"
	"majoo-test/repository/query"
)

type RoleUseCase struct {
	*Contract
}

func NewRoleUseCase(contract *Contract) usecase.IRoleUseCase {
	return &RoleUseCase{Contract: contract}
}

func (uc RoleUseCase) Browse(search string) (res []view_models.RoleVm, err error) {
	q := query.NewRoleQuery(uc.DB)
	roles, err := q.Browse(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productOutlet-store")
		return res, err
	}

	for _, role := range roles {
		res = append(res, view_models.NewRoleVm(role))
	}

	return res, nil
}
