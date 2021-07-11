package usecase

import "majoo-test/domain/view_models"

type ICategoryUseCase interface {
	Browse(search string)(res []view_models.CategoryVm,err error)
}
