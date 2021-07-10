package requests

type UserRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6"`
	MerchantID string `json:"merchant_id" validate:"required"`
	RoleID     int    `json:"role_id" validate:"required"`
}
