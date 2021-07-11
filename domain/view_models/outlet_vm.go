package view_models

import "majoo-test/domain/models"

type OutletVm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewOutletVm(model *models.Outlets) OutletVm {
	return OutletVm{
		ID:   model.Id(),
		Name: model.Name(),
	}
}
