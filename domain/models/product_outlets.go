package models

import (
	"database/sql"
	"time"
)

type ProductOutlets struct {
	id int64
	productID string
	outletID string
	stock int32
	price float64
	discount float64
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func (model *ProductOutlets) Id() int64 {
	return model.id
}

func (model *ProductOutlets) SetId(id int64) *ProductOutlets {
	model.id = id

	return model
}

func (model *ProductOutlets) ProductID() string {
	return model.productID
}

func (model *ProductOutlets) SetProductID(productID string)*ProductOutlets {
	model.productID = productID

	return model
}

func (model *ProductOutlets) OutletID() string {
	return model.outletID
}

func (model *ProductOutlets) SetOutletID(outletID string) *ProductOutlets{
	model.outletID = outletID

	return model
}

func (model *ProductOutlets) Stock() int32 {
	return model.stock
}

func (model *ProductOutlets) SetStock(stock int32) *ProductOutlets{
	model.stock = stock

	return model
}

func (model *ProductOutlets) Price() float64 {
	return model.price
}

func (model *ProductOutlets) SetPrice(price float64) *ProductOutlets{
	model.price = price

	return model
}

func (model *ProductOutlets) Discount() float64 {
	return model.discount
}

func (model *ProductOutlets) SetDiscount(discount float64)*ProductOutlets {
	model.discount = discount

	return model
}

func (model *ProductOutlets) CreatedAt() time.Time {
	return model.createdAt
}

func (model *ProductOutlets) SetCreatedAt(createdAt time.Time)*ProductOutlets {
	model.createdAt = createdAt

	return model
}

func (model *ProductOutlets) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *ProductOutlets) SetUpdatedAt(updatedAt time.Time) *ProductOutlets{
	model.updatedAt = updatedAt

	return model
}

func (model *ProductOutlets) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *ProductOutlets) SetDeletedAt(deletedAt sql.NullTime) *ProductOutlets{
	model.deletedAt = deletedAt

	return model
}


