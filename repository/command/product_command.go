package command

import (
	"database/sql"
	"majoo-test/domain/command"
	"majoo-test/domain/models"
)

type ProductCommand struct {
	DB *sql.DB
}

func NewProductCommand(DB *sql.DB) command.IProductCommand {
	return &ProductCommand{DB: DB}
}

func (ProductCommand) Add(model *models.Products, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO products (sku,name,unit_id,category_id,image_id,merchant_id,stock,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`
	err = tx.QueryRow(statement, model.Sku(), model.Name(), model.UnitId(), model.CategoryId(), model.ImageId(), model.MerchantId(), model.Stock(), model.CreatedAt(), model.UpdatedAt()).
		Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (ProductCommand) Edit(model *models.Products, tx *sql.Tx) (err error) {
	statement := `UPDATE products set sku=$1,name=$2,unit_id=$3,category_id=$4,image_id=$5,merchant_id=$6,updated_at=$7,stock=$8 WHERE id=$9`
	_, err = tx.Exec(statement, model.Sku(), model.Name(), model.UnitId(), model.CategoryId(), model.ImageId(), model.MerchantId(), model.UpdatedAt(), model.Stock(), model.Id())
	if err != nil {
		return err
	}

	return nil
}

func (ProductCommand) Delete(model *models.Products, tx *sql.Tx) (err error) {
	statement := `UPDATE products set updated_at=$1,deleted_at=$2 WHERE id=$3`
	_, err = tx.Exec(statement, model.UpdatedAt(), model.DeletedAt().Time, model.Id())
	if err != nil {
		return err
	}

	return nil
}
