package query

import "majoo-test/domain/models"

type IUnitQuery interface {
	Browse(search string) (res []*models.Units,err error)
}
