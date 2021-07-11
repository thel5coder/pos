package view_models

import (
	"majoo-test/pkg/str"
	"strings"
)

type OutletProductDetailsVm struct {
	ID            string          `json:"id"`
	Outlet OutletProductVm `json:"outlet"`
	Price         float64         `json:"price"`
}

func NewOutletDetailsVm(outletDetail string) (res []OutletProductDetailsVm) {
	if outletDetail != "" {
		outletDetails := str.Unique(strings.Split(outletDetail, ","))
		for _, outletDetail := range outletDetails {
			outletDetailArr := strings.Split(outletDetail, ":")
			res = append(res, build(outletDetailArr))
		}
	}

	return res
}

func build(outletDetail []string) OutletProductDetailsVm {
	return OutletProductDetailsVm{
		ID:            outletDetail[0],
		Outlet: NewOutletProductVm(outletDetail[1], outletDetail[2]),
		Price:         float64(str.StringToInt(outletDetail[3])),
	}
}
