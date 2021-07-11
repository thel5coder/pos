package view_models

type UserMerchantVm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserMerchantVm(id,name string) UserMerchantVm{
	return UserMerchantVm{
		ID:   id,
		Name: name,
	}
}
