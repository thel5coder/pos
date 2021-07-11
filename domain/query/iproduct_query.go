package query

import "majoo-test/domain/models"

type IProductQuery interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []*models.Products,err error)

	ReadBy(column,operator string,value interface{}) (res *models.Products,err error)

	CountBy(column,operator,ID string,value interface{}) (res int,err error)

	Count(search string) (res int,err error)
}
