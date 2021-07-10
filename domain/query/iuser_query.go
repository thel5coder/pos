package query

import "majoo-test/domain/models"

type IUserQuery interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []*models.Users, err error)

	ReadBy(column, operator string, value interface{}) (res *models.Users, err error)

	Count(search string) (res int, err error)

	CountBy(column, operator, ID string, value interface{}) (res int, err error)
}
