package command

import (
	"database/sql"
	"majoo-test/domain/command"
	"majoo-test/domain/models"
)

type ProductOutletCommand struct{
	DB *sql.DB
}



func NewProductOutletCommand(DB *sql.DB) command.IProductOutletCommand{
	return &ProductOutletCommand{DB: DB}
}

func (ProductOutletCommand) Add(model *models.ProductOutlets, tx *sql.Tx) (err error) {
	statement := `INSERT INTO product_outlets (product_id,outlet_id,price,created_at,updated_at) VALUES($1,$2,$3,$4,$5)`
	_,err = tx.Exec(statement,model.ProductID(),model.OutletID(),model.Price(),model.CreatedAt(),model.UpdatedAt())
	if err != nil {
		return err
	}

	return nil
}


func (ProductOutletCommand) DeleteByProduct(model *models.ProductOutlets, tx *sql.Tx) (err error) {
	statement := `DELETE FROM product_outlets WHERE product_id=$1`
	_,err = tx.Exec(statement,model.ProductID())
	if err != nil {
		return err
	}

	return nil
}

