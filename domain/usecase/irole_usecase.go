package usecase

import "majoo-test/domain/view_models"

type IRoleUseCase interface {
	Browse(search string) (res []view_models.RoleVm,err error)
}
