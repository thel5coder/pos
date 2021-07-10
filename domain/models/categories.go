package models

import (
	"database/sql"
	"time"
)

type Categories struct {
	id        string
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func (model *Categories) Id() string {
	return model.id
}

func (model *Categories) SetId(id string) *Categories {
	model.id = id

	return model
}

func (model *Categories) Name() string {
	return model.name
}

func (model *Categories) SetName(name string) *Categories {
	model.name = name

	return model
}

func (model *Categories) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Categories) SetCreatedAt(createdAt time.Time) *Categories {
	model.createdAt = createdAt

	return model
}

func (model *Categories) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Categories) SetUpdatedAt(updatedAt time.Time) *Categories {
	model.updatedAt = updatedAt

	return model
}

func (model *Categories) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Categories) SetDeletedAt(deletedAt sql.NullTime) *Categories {
	model.deletedAt = deletedAt

	return model
}
