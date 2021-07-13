package usecase

import "majoo-test/domain/view_models"

type IMerchantUseCase interface {
	Browse(search string) (res []view_models.UserMerchantVm,err error)
}
