package usecase

import (
	"majoo-test/domain/usecase"
	"majoo-test/domain/view_models"
	"majoo-test/pkg/functioncaller"
	"majoo-test/pkg/logruslogger"
	"majoo-test/repository/query"
)

type UnitUseCase struct {
	*Contract
}

func NewUnitUseCase(contract *Contract) usecase.IUnitUseCase {
	return &UnitUseCase{Contract: contract}
}

func (uc UnitUseCase) Browse(search string) (res []view_models.UnitVm, err error) {
	q := query.NewUnitQuery(uc.DB)

	units, err := q.Browse(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-unit-browse")
		return res, err
	}
	for _, unit := range units {
		res = append(res, view_models.NewUnitVm(unit))
	}

	return res, nil
}
