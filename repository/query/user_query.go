package query

import (
	"database/sql"
	"fmt"
	"majoo-test/domain/models"
	"majoo-test/domain/query"
	"strings"
)

type UserQuery struct {
	DB *sql.DB
}

func NewUserQuery(DB *sql.DB) query.IUserQuery {
	return &UserQuery{DB: DB}
}

func (q UserQuery) Browse(search, orderBy, sort string, limit, offset int) (res []*models.Users, err error) {
	statement := models.UserSelectStatement + ` ` + models.UserJoinStatement + ` ` + models.UserDefaultWhereStatement + ` AND (lower(u.email) like $1 OR lower(r.name) like $1) ORDER BY ` + orderBy + ` ` +
		sort + ` LIMIT $2 OFFSET $3`
	rows, err := q.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, err
	}

	model := models.NewUserModel()
	for rows.Next() {
		temp, err := model.ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}

	return res, nil
}

func (q UserQuery) ReadBy(column, operator string, value interface{}) (res *models.Users, err error) {
	statement := models.UserSelectStatement + ` ` + models.UserJoinStatement + ` ` + models.UserDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`
	row := q.DB.QueryRow(statement, value)
	model := models.NewUserModel()
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q UserQuery) Count(search string) (res int, err error) {
	statement := models.UserSelectCountStatement + ` ` + models.UserJoinStatement + ` ` + models.UserDefaultWhereStatement + ` AND (lower(u.email) like $1 OR lower(r.name) like $1)`
	err = q.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q UserQuery) CountBy(column, operator, ID string, value interface{}) (res int, err error) {
	statement := models.UserSelectCountStatement + ` ` + models.UserJoinStatement + ` ` + models.UserDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`
	countParams := []interface{}{value}
	if ID != "" {
		countParams = append(countParams, ID)
		statement += ` AND u.id<>$2`
	}
	fmt.Println(statement)

	err = q.DB.QueryRow(statement, countParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
