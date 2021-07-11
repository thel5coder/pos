package command

import (
	"database/sql"
	"majoo-test/domain/models"
)

type IProductOutletCommand interface {
	Add(model *models.ProductOutlets,tx *sql.Tx) (err error)


	DeleteByProduct(model *models.ProductOutlets,tx *sql.Tx) (err error)
}
