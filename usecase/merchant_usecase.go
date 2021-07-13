package usecase

import (
	"majoo-test/domain/usecase"
	"majoo-test/domain/view_models"
	"majoo-test/pkg/functioncaller"
	"majoo-test/pkg/logruslogger"
	"majoo-test/repository/query"
)

type MerchantUseCase struct {
	*Contract
}

func NewMerchantUseCase(contract *Contract) usecase.IMerchantUseCase {
	return &MerchantUseCase{Contract: contract}
}

func (uc MerchantUseCase) Browse(search string) (res []view_models.UserMerchantVm, err error) {
	q := query.NewMerchantQuery(uc.DB)
	merchants, err := q.Browse(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-merchant-browse")
		return res, err
	}
	for _, merchant := range merchants {
		res = append(res, view_models.NewUserMerchantVm(merchant.Id(), merchant.Name()))
	}

	return res, nil
}
