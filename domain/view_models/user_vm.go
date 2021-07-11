package view_models

import (
	"majoo-test/domain/models"
	"time"
)

type UserVm struct {
	ID        string         `json:"id"`
	Email     string         `json:"email"`
	Role      RoleVm         `json:"role"`
	Merchant  UserMerchantVm `json:"merchant"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

func NewUserVm(model *models.Users) UserVm {
	return UserVm{
		ID:        model.Id(),
		Email:     model.Email(),
		Role:      NewRoleVm(model.Roles),
		Merchant:  NewUserMerchantVm(model.MerchantID(), model.Merchants.Name()),
		CreatedAt: model.CreatedAt().Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt().Format(time.RFC3339),
	}
}
