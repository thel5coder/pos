package usecase

import (
	"majoo-test/domain/requests"
	"majoo-test/domain/view_models"
)

type IProductUseCase interface {
	Browse(search, orderBy, sort string, page, limit int) (res []view_models.ProductVm, pagination view_models.PaginationVm, err error)

	ReadBy(column, operator string, value interface{}) (res view_models.ProductVm, err error)

	Add(req *requests.ProductRequest) (res string, err error)

	Edit(req *requests.ProductRequest, ID string) (res string, err error)

	Delete(ID string) (err error)

	Count(search string) (res int, err error)

	CountBy(column, operator, ID string, value interface{}) (res int,err error)
}
