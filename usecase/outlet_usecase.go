package usecase

import (
	"majoo-test/domain/usecase"
	"majoo-test/domain/view_models"
	"majoo-test/pkg/functioncaller"
	"majoo-test/pkg/logruslogger"
	"majoo-test/repository/query"
)

type OutletUseCase struct {
	*Contract
}

func NewOutletUseCase(contract *Contract) usecase.IOutletUseCase {
	return &OutletUseCase{Contract: contract}
}

func (uc OutletUseCase) Browse(search string) (res []view_models.OutletVm, err error) {
	q := query.NewOutletQuery(uc.DB)

	outlets, err := q.Browse(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-unit-browse")
		return res, err
	}
	for _, outlet := range outlets {
		res = append(res, view_models.NewOutletVm(outlet))
	}

	return res, nil
}
