package command

import (
	"database/sql"
	"majoo-test/domain/command"
	"majoo-test/domain/models"
)

type UserCommand struct {
	DB *sql.DB
}

func NewUserCommand(DB *sql.DB) command.IUserCommand {
	return &UserCommand{DB: DB}
}

func (UserCommand) Add(model *models.Users, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO users(email,password,role_id,merchant_id,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING id`
	err = tx.QueryRow(statement, model.Email(), model.Password(), model.RoleID(),model.MerchantID(), model.CreatedAt(), model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (UserCommand) Edit(model *models.Users, tx *sql.Tx) (err error) {
	editParams := []interface{}{model.Email(), model.RoleID(),model.MerchantID(), model.UpdatedAt(), model.Id()}
	statement := `UPDATE users set email=$1,role_id$2,merchant_id=$3,updated_at$4 where id=$5`
	if model.Password() != "" {
		editParams = append(editParams, model.Password())
		statement =`UPDATE users set email=$1,role_id$2,merchant_id=$3,updated_at$4,password=$6 where id=$5`
	}

	_, err = tx.Exec(statement, editParams...)
	if err != nil {
		return err
	}

	return nil
}

func (UserCommand) Delete(model *models.Users, tx *sql.Tx) (err error) {
	statement := `UPDATE users set updated_at=$1,deleted_at=$2 where id=$3`

	_, err = tx.Exec(statement, model.UpdatedAt(), model.DeletedAt().Time, model.Id())
	if err != nil {
		return err
	}

	return nil
}
