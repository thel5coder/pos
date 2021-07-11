package usecase

import (
	"majoo-test/domain/usecase"
	"majoo-test/domain/view_models"
	"majoo-test/pkg/functioncaller"
	"majoo-test/pkg/logruslogger"
	"majoo-test/repository/query"
)

type CategoryUseCase struct {
	*Contract
}

func NewCategoryUseCase(contract *Contract) usecase.ICategoryUseCase {
	return &CategoryUseCase{Contract: contract}
}

func (uc CategoryUseCase) Browse(search string) (res []view_models.CategoryVm, err error) {
	q := query.NewCategoryQuery(uc.DB)

	categories, err := q.Browse(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-unit-browse")
		return res, err
	}
	for _, category := range categories {
		res = append(res, view_models.NewCategoryVm(category))
	}

	return res, nil
}
