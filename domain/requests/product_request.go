package requests

type ProductRequest struct {
	SKU           string                 `json:"sku" validate:"required"`
	Name          string                 `json:"name" validate:"required"`
	UnitID        int64                  `json:"unit_id" validate:"required"`
	CategoryID    string                 `json:"category_id" validate:"required"`
	ImageID       string                 `json:"image_id"`
	OutletDetails []ProductOutletRequest `json:"outlet_details"`
}
