package query

import "majoo-test/domain/models"

type IOutletQuery interface {
	Browse(search string) (res []*models.Outlets,err error)
}
