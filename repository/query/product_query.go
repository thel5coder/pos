package query

import (
	"database/sql"
	"majoo-test/domain/models"
	"majoo-test/domain/query"
	"strings"
)

type ProductQuery struct {
	DB *sql.DB
}

func NewProductQuery(DB *sql.DB) query.IProductQuery {
	return &ProductQuery{DB: DB}
}

func (q ProductQuery) Browse(search, orderBy, sort string, limit, offset int) (res []*models.Products, err error) {
	statement := models.ProductSelectStatement + ` ` + models.ProductJoinStatement + ` ` + models.ProductDefaultWhereStatement + ` AND (LOWER(p.name) LIKE $1 OR LOWER(p.sku) LIKE $1) ` +
		models.ProductGroupByStatement + ` ` + `ORDER BY ` + orderBy + ` ` + sort + ` LIMIT $2 OFFSET $3`
	rows, err := q.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewProductModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}

	return res, nil
}

func (q ProductQuery) ReadBy(column, operator string, value interface{}) (res *models.Products, err error) {
	statement := models.ProductSelectStatement + ` ` + models.ProductJoinStatement + ` ` + models.UserDefaultWhereStatement + ` AND ` + column + `` + operator + `$1 ` + models.ProductGroupByStatement
	row := q.DB.QueryRow(statement, value)
	res, err = models.NewProductModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q ProductQuery) CountBy(column, operator,ID string, value interface{}) (res int, err error) {
	whereStatement := models.ProductDefaultWhereStatement + ` AND ` + column + `` + operator + `$1 `
	countParams := []interface{}{value}
	if ID != "" {
		whereStatement += `AND p.id<>$2 `
		countParams = append(countParams,ID)
	}

	statement := models.ProductSelectCountStatement + ` ` + models.ProductJoinStatement + ` ` + whereStatement
	err = q.DB.QueryRow(statement, countParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q ProductQuery) Count(search string) (res int, err error) {
	statement := models.ProductSelectCountStatement + ` ` + models.ProductJoinStatement + ` ` + models.ProductDefaultWhereStatement + ` AND (LOWER(p.name) LIKE $1 OR LOWER(p.sku) LIKE $1) ` +
		models.ProductGroupByStatement
	err = q.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
