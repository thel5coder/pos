package requests

type ProductOutletRequest struct {
	OutletID string `json:"outlet_id" validate:"required"`
	Price    float64  `json:"price"`
}
