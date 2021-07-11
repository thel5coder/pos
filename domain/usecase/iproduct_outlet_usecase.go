package usecase

import "majoo-test/domain/requests"

type IProductOutletUseCase interface {
	Add(req *requests.ProductOutletRequest,productID string) (err error)

	DeleteByProduct(productID string) (err error)

	CountByProduct(productID string) (res int,err error)

	Store(req []requests.ProductOutletRequest,productID string)(err error)
}
