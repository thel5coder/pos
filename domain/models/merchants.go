package models

import (
	"database/sql"
	"time"
)

type Merchants struct {
	id        string
	userId    string
	name      string
	logo      sql.NullString
	address   sql.NullString
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func NewMerchantsModel() *Merchants {
	return &Merchants{}
}

func (model *Merchants) Id() string {
	return model.id
}

func (model *Merchants) SetId(id string) *Merchants {
	model.id = id

	return model
}

func (model *Merchants) UserId() string {
	return model.userId
}

func (model *Merchants) SetUserId(userId string) *Merchants {
	model.userId = userId

	return model
}

func (model *Merchants) Name() string {
	return model.name
}

func (model *Merchants) SetName(name string) *Merchants {
	model.name = name

	return model
}

func (model *Merchants) Logo() sql.NullString {
	return model.logo
}

func (model *Merchants) SetLogo(logo sql.NullString) *Merchants {
	model.logo = logo

	return model
}

func (model *Merchants) Address() sql.NullString {
	return model.address
}

func (model *Merchants) SetAddress(address sql.NullString) *Merchants {
	model.address = address

	return model
}

func (model *Merchants) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Merchants) SetCreatedAt(createdAt time.Time) *Merchants {
	model.createdAt = createdAt

	return model
}

func (model *Merchants) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Merchants) SetUpdatedAt(updatedAt time.Time) *Merchants {
	model.updatedAt = updatedAt

	return model
}

func (model *Merchants) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Merchants) SetDeletedAt(deletedAt sql.NullTime) *Merchants {
	model.deletedAt = deletedAt

	return model
}

const (
	MerchantSelectStatement = `SELECT id,name FROM merchants`
)

func (model *Merchants) ScanRows(rows *sql.Rows) (*Merchants, error) {
	err := rows.Scan(&model.id, &model.name)
	if err != nil {
		return model, err
	}

	return model, err
}
