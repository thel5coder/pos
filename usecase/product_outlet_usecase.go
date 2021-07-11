package usecase

import (
	"majoo-test/domain/models"
	"majoo-test/domain/requests"
	"majoo-test/domain/usecase"
	"majoo-test/pkg/functioncaller"
	"majoo-test/pkg/logruslogger"
	"majoo-test/repository/command"
	"majoo-test/repository/query"
)

type ProductOutletUseCase struct {
	*Contract
}

func NewProductOutletUseCase(contract *Contract) usecase.IProductOutletUseCase {
	return &ProductOutletUseCase{Contract: contract}
}

func (uc ProductOutletUseCase) Add(req *requests.ProductOutletRequest, productID string) (err error) {
	cmd := command.NewProductOutletCommand(uc.DB)

	model := models.NewProductOutletModel().SetOutletID(req.OutletID).SetProductID(productID).SetPrice(req.Price)
	err = cmd.Add(model, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-productOutlet-add")
		return err
	}

	return nil
}

func (uc ProductOutletUseCase) DeleteByProduct(productID string) (err error) {
	cmd := command.NewProductOutletCommand(uc.DB)

	model := models.NewProductOutletModel().SetProductID(productID)
	err = cmd.DeleteByProduct(model,uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-productOutlet-deleteByProduct")
		return err
	}

	return nil
}

func (uc ProductOutletUseCase) CountByProduct(productID string) (res int, err error) {
	q := query.NewProductOutletQuery(uc.DB)
	res, err = q.CountByProduct(productID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-productOutlet-countByProductID")
		return res, err
	}

	return res, nil
}

func (uc ProductOutletUseCase) Store(req []requests.ProductOutletRequest, productID string) (err error) {
	count,err := uc.CountByProduct(productID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productOutlet-countByProduct")
		return err
	}
	if count > 0 {
		err = uc.DeleteByProduct(productID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productOutlet-deleteByProduct")
			return err
		}
	}

	for _,productOutlet := range req {
		err = uc.Add(&productOutlet,productID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productOutlet-add")
			return err
		}
	}

	return nil
}
