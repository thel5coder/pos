package query

import (
	"database/sql"
	"majoo-test/domain/models"
	"majoo-test/domain/query"
	"strings"
)

type OutletQuery struct {
	DB *sql.DB
}

func NewOutletQuery(DB *sql.DB) query.IOutletQuery {
	return &OutletQuery{DB: DB}
}

func (q OutletQuery) Browse(search string) (res []*models.Outlets, err error) {
	statement := models.OutletSelectStatement + ` WHERE LOWER(name) LIKE $1`
	rows, err := q.DB.Query(statement, "%"+strings.ToLower(search)+"%")
	if err != nil {
		return res, err
	}
	for rows.Next() {
		temp, err := models.NewOutletModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}

	return res, nil
}
