package query

import (
	"database/sql"
	"majoo-test/domain/models"
	"majoo-test/domain/query"
	"strings"
)

type CategoryQuery struct {
	DB *sql.DB
}

func NewCategoryQuery(DB *sql.DB) query.ICategoryQuery {
	return &CategoryQuery{DB: DB}
}

func (q CategoryQuery) Browse(search string) (res []*models.Categories, err error) {
	statement := models.CategorySelectStatement + ` WHERE LOWER(name) LIKE $1`
	rows, err := q.DB.Query(statement, "%"+strings.ToLower(search)+"%")
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewCategoryModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}

	return res, nil
}
