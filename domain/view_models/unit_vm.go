package view_models

import "majoo-test/domain/models"

type UnitVm struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewUnitVm(model *models.Units) UnitVm {
	return UnitVm{
		ID:   model.Id(),
		Name: model.Name(),
	}
}
