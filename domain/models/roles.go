package models

import "database/sql"

type Roles struct {
	id   int64
	name string
	slug string
}

func NewRolesModel() *Roles{
	return &Roles{}
}

func (model *Roles) Id() int64 {
	return model.id
}

func (model *Roles) SetId(id int64) *Roles {
	model.id = id

	return model
}

func (model *Roles) Name() string {
	return model.name
}

func (model *Roles) SetName(name string) *Roles {
	model.name = name

	return model
}

func (model *Roles) Slug() string {
	return model.slug
}

func (model *Roles) SetSlug(slug string) *Roles {
	model.slug = slug

	return model
}

const (
	RoleSelectStatement = `SELECT id,name FROM roles`
)

func(model *Roles) ScanRows(rows *sql.Rows) (*Roles,error){
	err := rows.Scan(&model.id,&model.name)
	if err != nil {
		return model, err
	}

	return model, nil
}
