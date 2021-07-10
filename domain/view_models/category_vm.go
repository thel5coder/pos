package view_models

import "majoo-test/domain/models"

type CategoryVm struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func NewCategoryVm(model *models.Categories) CategoryVm{
	return CategoryVm{
		ID:   model.Id(),
		Name: model.Name(),
	}
}
