package query

import "majoo-test/domain/models"

type IMerchantQuery interface {
	Browse(search string) (res []*models.Merchants,err error)
}
