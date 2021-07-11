package command

import (
	"database/sql"
	"majoo-test/domain/models"
)

type IProductCommand interface {
	Add(model *models.Products, tx *sql.Tx) (res string, err error)

	Edit(model *models.Products, tx *sql.Tx) (err error)

	Delete(model *models.Products, tx *sql.Tx) (err error)
}
