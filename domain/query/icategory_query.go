package query

import "majoo-test/domain/models"

type ICategoryQuery interface {
	Browse(search string) (res []*models.Categories, err error)
}
