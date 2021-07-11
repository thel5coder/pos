package query

import (
	"database/sql"
	"majoo-test/domain/models"
	"majoo-test/domain/query"
	"strings"
)

type RoleQuery struct{
	DB *sql.DB
}

func NewRoleQuery(DB *sql.DB) query.IRoleQuery{
	return &RoleQuery{DB: DB}
}

func (q RoleQuery) Browse(search string) (res []*models.Roles, err error) {
	statement := models.RoleSelectStatement+` WHERE LOWER(name) like $1`
	rows,err := q.DB.Query(statement,"%"+strings.ToLower(search)+"%")
	if err != nil {
		return res,err
	}

	for rows.Next(){
		temp,err := models.NewRolesModel().ScanRows(rows)
		if err != nil {
			return res,err
		}
		res = append(res,temp)
	}

	return res,nil
}

