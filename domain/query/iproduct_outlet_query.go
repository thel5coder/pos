package query

type IProductOutletQuery interface {
	CountByProduct(productID string) (res int,err error)
}
