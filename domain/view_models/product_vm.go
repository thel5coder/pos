package view_models

import (
	"majoo-test/domain/models"
	"time"
)

type ProductVm struct {
	ID        string     `json:"id"`
	SKU       string     `json:"sku"`
	Name      string     `json:"name"`
	Unit      UnitVm     `json:"unit"`
	Category  CategoryVm `json:"category"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

func NewProductVm(model *models.Products) ProductVm{
	return ProductVm{
		ID:        model.Id(),
		SKU:       model.Sku(),
		Name:      model.Name(),
		Unit:      NewUnitVm(model.Unit),
		Category:  NewCategoryVm(model.Category),
		CreatedAt: model.CreatedAt().Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt().Format(time.RFC3339),
	}
}
