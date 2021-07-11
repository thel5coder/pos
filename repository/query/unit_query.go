package query

import (
	"database/sql"
	"majoo-test/domain/models"
	"majoo-test/domain/query"
	"strings"
)

type UnitQuery struct {
	DB *sql.DB
}

func NewUnitQuery(DB *sql.DB) query.IUnitQuery {
	return &UnitQuery{DB: DB}
}

func (q UnitQuery) Browse(search string) (res []*models.Units, err error) {
	statement := models.UnitSelectStatement + ` WHERE LOWER(name) LIKE $1`
	rows, err := q.DB.Query(statement, "%"+strings.ToLower(search)+"%")
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewUnitModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}

	return res, nil
}
