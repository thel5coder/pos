package requests

type ProductRequest struct {
	SKU           string                 `json:"sku"`
	Name          string                 `json:"name"`
	UnitID        int64                  `json:"unit_id"`
	CategoryID    string                 `json:"category_id"`
	ImageID       string                 `json:"image_id"`
	OutletDetails []ProductOutletRequest `json:"outlet_details"`
}
