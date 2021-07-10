package models

import (
	"database/sql"
	"time"
)

type Units struct {
	id        int64
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func (model *Units) Id() int64 {
	return model.id
}

func (model *Units) SetId(id int64) *Units{
	model.id = id

	return model
}

func (model *Units) Name() string {
	return model.name
}

func (model *Units) SetName(name string) *Units{
	model.name = name

	return model
}

func (model *Units) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Units) SetCreatedAt(createdAt time.Time) *Units{
	model.createdAt = createdAt

	return model
}

func (model *Units) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Units) SetUpdatedAt(updatedAt time.Time) *Units{
	model.updatedAt = updatedAt

	return model
}

func (model *Units) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Units) SetDeletedAt(deletedAt sql.NullTime) *Units{
	model.deletedAt = deletedAt

	return model
}
