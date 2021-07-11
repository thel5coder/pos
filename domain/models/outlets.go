package models

import (
	"database/sql"
	"time"
)

type Outlets struct {
	id         string
	merchantId string
	name       string
	address    string
	createdAt  time.Time
	updatedAt  time.Time
	deletedAt  sql.NullTime
}

func NewOutletModel() *Outlets {
	return &Outlets{}
}

func (model *Outlets) Id() string {
	return model.id
}

func (model *Outlets) SetId(id string) *Outlets {
	model.id = id

	return model
}

func (model *Outlets) MerchantId() string {
	return model.merchantId
}

func (model *Outlets) SetMerchantId(merchantId string) *Outlets {
	model.merchantId = merchantId

	return model
}

func (model *Outlets) Name() string {
	return model.name
}

func (model *Outlets) SetName(name string) *Outlets {
	model.name = name

	return model
}

func (model *Outlets) Address() string {
	return model.address
}

func (model *Outlets) SetAddress(address string) *Outlets {
	model.address = address

	return model
}

func (model *Outlets) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Outlets) SetCreatedAt(createdAt time.Time) *Outlets {
	model.createdAt = createdAt

	return model
}

func (model *Outlets) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Outlets) SetUpdatedAt(updatedAt time.Time) *Outlets {
	model.updatedAt = updatedAt

	return model
}

func (model *Outlets) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Outlets) SetDeletedAt(deletedAt sql.NullTime) *Outlets {
	model.deletedAt = deletedAt

	return model
}

const (
	OutletSelectStatement = `SELECT id,name FROM outlets`
)

func (model *Outlets) ScanRows(rows *sql.Rows) (*Outlets, error) {
	err := rows.Scan(&model.id, &model.name)
	if err != nil {
		return model, err
	}

	return model, nil
}
