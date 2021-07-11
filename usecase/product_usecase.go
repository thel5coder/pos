package usecase

import (
	"errors"
	"fmt"
	"majoo-test/domain/models"
	"majoo-test/domain/requests"
	"majoo-test/domain/usecase"
	"majoo-test/domain/view_models"
	"majoo-test/pkg/functioncaller"
	"majoo-test/pkg/logruslogger"
	"majoo-test/pkg/messages"
	"majoo-test/repository/command"
	"majoo-test/repository/query"
	"time"
)

type ProductUseCase struct {
	*Contract
}

func NewProductUseCase(contract *Contract) usecase.IProductUseCase {
	return &ProductUseCase{Contract: contract}
}

func (uc ProductUseCase) Browse(search, orderBy, sort string, page, limit int) (res []view_models.ProductVm, pagination view_models.PaginationVm, err error) {
	q := query.NewProductQuery(uc.DB)
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	products, err := q.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-browse")
		return res, pagination, err
	}
	for _, product := range products {
		res = append(res, view_models.NewProductVm(product))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-product-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc ProductUseCase) ReadBy(column, operator string, value interface{}) (res view_models.ProductVm, err error) {
	q := query.NewProductQuery(uc.DB)

	product, err := q.ReadBy(column, operator, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-readBy")
		return res, err
	}
	res = view_models.NewProductVm(product)

	return res, nil
}

func (uc ProductUseCase) Add(req *requests.ProductRequest) (res string, err error) {
	cmd := command.NewProductCommand(uc.DB)
	now := time.Now().UTC()

	count, err := uc.CountBy("p.sku", "=", "", req.SKU)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-product-countBySKU")
		return res, err
	}
	if count > 0 {
		return res, errors.New(messages.DataAlreadyExist)
	}

	fmt.Println(uc.MerchantID)
	model := models.NewProductModel().SetSku(req.SKU).SetName(req.Name).SetUnitId(req.UnitID).SetCategoryId(req.CategoryID).SetImageId(req.ImageID).
		SetCreatedAt(now).SetUpdatedAt(now).SetMerchantId(uc.MerchantID)
	res, err = cmd.Add(model, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-add")
		return res, err
	}

	productOutletUc := NewProductOutletUseCase(uc.Contract)
	err = productOutletUc.Store(req.OutletDetails, res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productOutlet-store")
		return res, err
	}

	return res, nil
}

func (uc ProductUseCase) Edit(req *requests.ProductRequest, ID string) (res string, err error) {
	cmd := command.NewProductCommand(uc.DB)
	now := time.Now().UTC()

	count, err := uc.CountBy("p.sku", "=", ID, req.SKU)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-product-countBySKU")
		return res, err
	}
	if count > 0 {
		return res, errors.New(messages.DataAlreadyExist)
	}

	model := models.NewProductModel().SetSku(req.SKU).SetName(req.Name).SetUnitId(req.UnitID).SetCategoryId(req.CategoryID).SetImageId(req.ImageID).
		SetCreatedAt(now).SetUpdatedAt(now).SetId(ID).SetMerchantId(uc.MerchantID)
	err = cmd.Edit(model, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-add")
		return res, err
	}

	productOutletUc := NewProductOutletUseCase(uc.Contract)
	err = productOutletUc.Store(req.OutletDetails, ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productOutlet-store")
		return res, err
	}
	res = ID

	return res, nil
}

func (uc ProductUseCase) Delete(ID string) (err error) {
	cmd := command.NewProductCommand(uc.DB)
	now := time.Now().UTC()

	model := models.NewProductModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(ID)
	err = cmd.Delete(model, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-delete")
		return err
	}

	productOutletUc := NewProductOutletUseCase(uc.Contract)
	err = productOutletUc.Store(nil, ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productOutlet-store")
		return err
	}

	return nil
}

func (uc ProductUseCase) Count(search string) (res int, err error) {
	q := query.NewProductQuery(uc.DB)
	res, err = q.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-count")
		return res, err
	}

	return res, nil
}

func (uc ProductUseCase) CountBy(column, operator, ID string, value interface{}) (res int, err error) {
	q := query.NewProductQuery(uc.DB)
	res, err = q.CountBy(column, operator, ID, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-countBy")
		return res, err
	}

	return res, nil
}
