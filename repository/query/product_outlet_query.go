package query

import (
	"database/sql"
	"majoo-test/domain/query"
)

type ProductOutletQuery struct{
	DB *sql.DB
}

func NewProductOutletQuery(DB *sql.DB) query.IProductOutletQuery{
	return &ProductOutletQuery{DB: DB}
}

func (q ProductOutletQuery) CountByProduct(productID string) (res int, err error) {
	statement := `SELECT COUNT(id) FROM product_outlets WHERE product_id=$1`
	err = q.DB.QueryRow(statement,productID).Scan(&res)
	if err != nil {
		return res,err
	}

	return res,err
}

