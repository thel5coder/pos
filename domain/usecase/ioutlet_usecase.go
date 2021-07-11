package usecase

import "majoo-test/domain/view_models"

type IOutletUseCase interface {
	Browse(search string)(res []view_models.OutletVm,err error)
}
