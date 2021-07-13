package query

import (
	"database/sql"
	"majoo-test/domain/models"
	"majoo-test/domain/query"
	"strings"
)

type MerchantQuery struct {
	DB *sql.DB
}

func NewMerchantQuery(DB *sql.DB) query.IMerchantQuery {
	return &MerchantQuery{DB: DB}
}

func (q MerchantQuery) Browse(search string) (res []*models.Merchants, err error) {
	statement := models.MerchantSelectStatement + ` WHERE LOWER(name) LIKE $1`
	rows, err := q.DB.Query(statement, "%"+strings.ToLower(search)+"%")
	if err != nil {
		return res, err
	}
	for rows.Next() {
		temp, err := models.NewMerchantsModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}

	return res, nil
}
