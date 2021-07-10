package view_models

import "majoo-test/domain/models"

type RoleVm struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func NewRoleVm(model *models.Roles) RoleVm {
	return RoleVm{
		ID:   model.Id(),
		Name: model.Name(),
		Slug: model.Slug(),
	}
}
