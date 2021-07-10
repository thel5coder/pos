package command

import (
	"database/sql"
	"majoo-test/domain/models"
)

type IUserCommand interface {
	Add(model *models.Users, tx *sql.Tx) (res string, err error)

	Edit(model *models.Users, tx *sql.Tx) (err error)

	Delete(model *models.Users, tx *sql.Tx) (err error)
}
