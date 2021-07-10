package requests

type ProductOutletRequest struct {
	OutletID string  `json:"outlet_id"`
	Stock    int64   `json:"stock"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"`
}
