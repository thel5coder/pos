package usecase

import (
	"errors"
	"fmt"
	"majoo-test/domain/models"
	"majoo-test/domain/requests"
	"majoo-test/domain/usecase"
	"majoo-test/domain/view_models"
	"majoo-test/pkg/functioncaller"
	"majoo-test/pkg/hashing"
	"majoo-test/pkg/logruslogger"
	"majoo-test/pkg/messages"
	"majoo-test/repository/command"
	"majoo-test/repository/query"
	"time"
)

type UserUseCase struct {
	*Contract
}

func NewUserUseCase(contract *Contract) usecase.IUserUseCase {
	return &UserUseCase{Contract: contract}
}

func (uc UserUseCase) Browse(search, orderBy, sort string, page, limit int) (res []view_models.UserVm, pagination view_models.PaginationVm, err error) {
	q := query.NewUserQuery(uc.DB)
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	users, err := q.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-browse")
		return res, pagination, err
	}
	for _, user := range users {
		res = append(res, view_models.NewUserVm(user))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc UserUseCase) ReadBy(column, operator string, value interface{}) (res view_models.UserVm, err error) {
	q := query.NewUserQuery(uc.DB)

	user, err := q.ReadBy(column, operator, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-readBy")
		return res, err
	}
	res = view_models.NewUserVm(user)

	return res, nil
}

func (uc UserUseCase) Add(req *requests.UserRequest) (res string, err error) {
	cmd := command.NewUserCommand(uc.DB)
	now := time.Now().UTC()

	count, err := uc.CountBy("u.email", "=", "", req.Email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-countByEmail")
		return res, err
	}
	fmt.Println(count)
	if count > 0 {
		return res, errors.New(messages.DataAlreadyExist)
	}
	hashed, _ := hashing.HashAndSalt(req.Password)
	model := models.NewUserModel().SetEmail(req.Email).SetRoleID(req.RoleID).SetMerchantID(req.MerchantID).SetPassword(hashed).SetCreatedAt(now).SetUpdatedAt(now)
	res, err = cmd.Add(model, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-add")
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) Edit(req *requests.UserRequest, ID string) (res string, err error) {
	cmd := command.NewUserCommand(uc.DB)
	now := time.Now().UTC()

	count, err := uc.CountBy("u.email", "=", ID, req.Email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-countByEmail")
		return res, err
	}
	if count > 0 {
		return res, errors.New(messages.DataAlreadyExist)
	}

	hashed, _ := hashing.HashAndSalt(req.Password)
	model := models.NewUserModel().SetEmail(req.Email).SetRoleID(req.RoleID).SetMerchantID(req.MerchantID).SetPassword(hashed).SetCreatedAt(now).SetUpdatedAt(now)
	err = cmd.Edit(model, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-edit")
		return res, err
	}
	res = ID

	return res, nil
}

func (uc UserUseCase) Delete(ID string) (err error) {
	cmd := command.NewUserCommand(uc.DB)
	now := time.Now().UTC()

	model := models.NewUserModel().SetUpdatedAt(now).SetDeletedAt(now)
	err = cmd.Delete(model, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-delete")
		return err
	}

	return nil
}

func (uc UserUseCase) Count(search string) (res int, err error) {
	repository := query.NewUserQuery(uc.DB)
	res, err = repository.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-count")
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) CountBy(column, operator, ID string, value interface{}) (res int, err error) {
	repository := query.NewUserQuery(uc.DB)
	res, err = repository.CountBy(column, operator, ID, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-countBy")
		return res, err
	}

	return res, nil
}
