package query

import "majoo-test/domain/models"

type IRoleQuery interface {
	Browse(search string) (res []*models.Roles,err error)
}
