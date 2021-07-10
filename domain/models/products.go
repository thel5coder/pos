package models

import (
	"database/sql"
	"time"
)

type Products struct {
	id         string
	sku        string
	name       string
	unitId     int64
	categoryId string
	imageId    string
	merchantId string

	Unit     *Units
	Category *Categories
	Merchant Merchants

	OutletDetails sql.NullString

	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func (model *Products) ImageId() string {
	return model.imageId
}

func (model *Products) SetImageId(imageId string) *Products {
	model.imageId = imageId

	return model
}

func (model *Products) Id() string {
	return model.id
}

func (model *Products) SetId(id string) *Products {
	model.id = id

	return model
}

func (model *Products) Sku() string {
	return model.sku
}

func (model *Products) SetSku(sku string) *Products {
	model.sku = sku

	return model
}

func (model *Products) Name() string {
	return model.name
}

func (model *Products) SetName(name string) *Products {
	model.name = name

	return model
}

func (model *Products) UnitId() int64 {
	return model.unitId
}

func (model *Products) SetUnitId(unitId int64) *Products {
	model.unitId = unitId

	return model
}

func (model *Products) CategoryId() string {
	return model.categoryId
}

func (model *Products) SetCategoryId(categoryId string) *Products {
	model.categoryId = categoryId

	return model
}

func (model *Products) MerchantId() string {
	return model.merchantId
}

func (model *Products) SetMerchantId(merchantId string) *Products {
	model.merchantId = merchantId

	return model
}

func (model *Products) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Products) SetCreatedAt(createdAt time.Time) *Products {
	model.createdAt = createdAt

	return model
}

func (model *Products) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Products) SetUpdatedAt(updatedAt time.Time) *Products {
	model.updatedAt = updatedAt

	return model
}

func (model *Products) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Products) SetDeletedAt(deletedAt sql.NullTime) *Products {
	model.deletedAt = deletedAt

	return model
}

const (
	ProductListSelectStatement = `SELECT p.id,p.sku,p.name,p.unit_id,p.category_id,p.image_id,p.merchant_id,p.created_at,p.updated_at,` +
		`ARRAY_TO_STRING(ARRAY_AGG(po.id ||':'|| o.id ||':'|| o.name||':'|| po.stock ||':'|| po.price ||':'|| COALESCE(po.discount,0)),','),c.id,c.name,u.id,u.name FROM products p`
	ProductJoinStatement = `LEFT JOIN product_outlets po ON po.product_id = p.id AND po.deleted_at IS NULL` +
		`LEFT JOIN outlets o ON o.id = po.outlet_id AND o.deleted_at IS NULL` +
		`INNER JOIN units u ON u.id=p.unit_id AND u.deleted_at IS NULL` +
		`INNER JOIN categories c on c.id=p.category_id AND c.deleted_at IS NULL` +
		`INNER JOIN merchants m ON m.id = p.merchant_id AND m.deleted_at IS NULL` +
		`INNER JOIN users us ON us.id = m.user_id AND u.deleted_at IS NULL`
	ProductDefaultWhereStatement = `WHERE p.deleted_at IS NULL`
	ProductGroupByStatement      = `GROUP BY p.id,u.id,c.id,m.id,us.id`
)

func (model *Products) ScanRows(rows *sql.Rows) (*Products, error) {
	err := rows.Scan(&model.id, &model.sku, &model.name, &model.unitId, &model.categoryId, &model.imageId, &model.merchantId, &model.createdAt, &model.updatedAt, &model.OutletDetails,
		&model.Category.id, &model.Category.name, &model.Unit.id, &model.Unit.name)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Products) ScanRow(row *sql.Row) (*Products, error) {
	err := row.Scan(&model.id, &model.sku, &model.name, &model.unitId, &model.categoryId, &model.imageId, &model.merchantId, &model.createdAt, &model.updatedAt, &model.OutletDetails,
		&model.Category.id, &model.Category.name, &model.Unit.id, &model.Unit.name)
	if err != nil {
		return model, err
	}

	return model, nil
}
