package view_models

type OutletProductVm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewOutletProductVm(id,name string) OutletProductVm{
	return OutletProductVm{
		ID:   id,
		Name: name,
	}
}
