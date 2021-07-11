package usecase

import "majoo-test/domain/view_models"

type IUnitUseCase interface {
	Browse(search string)(res []view_models.UnitVm,err error)
}
