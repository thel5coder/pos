package models

import (
	"database/sql"
	"time"
)

type Users struct {
	id         string
	email      string
	password   string
	roleID     int
	Roles      *Roles
	merchantID string
	Merchants  *Merchants
	createdAt  time.Time
	updatedAt  time.Time
	deletedAt  sql.NullTime
}


func (model *Users) MerchantID() string {
	return model.merchantID
}

func (model *Users) SetMerchantID(merchantID string) *Users {
	model.merchantID = merchantID

	return model
}

func NewUserModel() *Users {
	return &Users{}
}

func (model *Users) Id() string {
	return model.id
}

func (model *Users) SetId(id string) *Users {
	model.id = id

	return model
}

func (model *Users) Email() string {
	return model.email
}

func (model *Users) SetEmail(email string) *Users {
	model.email = email

	return model
}

func (model *Users) Password() string {
	return model.password
}

func (model *Users) SetPassword(password string) *Users {
	model.password = password

	return model
}

func (model *Users) RoleID() int {
	return model.roleID
}

func (model *Users) SetRoleID(roleID int) *Users {
	model.roleID = roleID

	return model
}

func (model *Users) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Users) SetCreatedAt(createdAt time.Time) *Users {
	model.createdAt = createdAt

	return model
}

func (model *Users) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Users) SetUpdatedAt(updatedAt time.Time) *Users {
	model.updatedAt = updatedAt

	return model
}

func (model *Users) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Users) SetDeletedAt(deletedAt time.Time) *Users {
	model.deletedAt.Time = deletedAt

	return model
}

const (
	UserSelectStatement      = `SELECT u.id,u.email,u.password,u.role_id,u.created_at,u.updated_at,r.id,r.name,r.slug,m.id,m.name FROM users u `
	UserSelectCountStatement = `SELECT COUNT(u.id) FROM users u `
	UserJoinStatement        = `INNER JOIN roles r on r.id=u.role_id ` +
		`INNER JOIN merchants m ON m.id = u.merchant_id AND m.deleted_at IS NULL `
	UserDefaultWhereStatement = `WHERE u.deleted_at IS NULL`
)

func (model *Users) ScanRows(rows *sql.Rows) (*Users, error) {
	model.Roles = NewRolesModel()
	model.Merchants = NewMerchantsModel()
	err := rows.Scan(&model.id, &model.email, &model.password, &model.roleID, &model.createdAt, &model.updatedAt, &model.Roles.id, &model.Roles.name, &model.Roles.slug,
		&model.merchantID, &model.Merchants.name)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Users) ScanRow(row *sql.Row) (*Users, error) {
	model.Roles = NewRolesModel()
	model.Merchants = NewMerchantsModel()
	err := row.Scan(&model.id, &model.email, &model.password, &model.roleID, &model.createdAt, &model.updatedAt, &model.Roles.id, &model.Roles.name, &model.Roles.slug,
		&model.merchantID, &model.Merchants.name)
	if err != nil {
		return model, err
	}

	return model, nil
}
