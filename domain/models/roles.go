package models

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
