package view_models

import (
	"majoo-test/pkg/str"
	"strings"
)

type OutletProductDetailsVm struct {
	ID            string          `json:"id"`
	OutletProduct OutletProductVm `json:"outlet_product"`
	Stock         int32           `json:"stock"`
	Price         float64         `json:"price"`
	Discount      float64         `json:"discount"`
}

func NewOutletDetailsVm(outletDetail string) (res []OutletProductDetailsVm) {
	outletDetails := str.Unique(strings.Split(outletDetail, ","))
	for _, outletDetail := range outletDetails {
		outletDetailArr := strings.Split(outletDetail, ":")
		res = append(res, build(outletDetailArr))
	}

	return res
}

func build(outletDetail []string) OutletProductDetailsVm {
	return OutletProductDetailsVm{
		ID:            outletDetail[0],
		OutletProduct: NewOutletProductVm(outletDetail[1], outletDetail[2]),
		Stock:         int32(str.StringToInt(outletDetail[3])),
		Price:         float64(str.StringToInt(outletDetail[4])),
		Discount:      float64(str.StringToInt(outletDetail[5])),
	}
}
