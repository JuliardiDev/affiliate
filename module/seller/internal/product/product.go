package product

import "context"

var mock bool

type (
	IProduct interface {
		FetchAllDB(limit int64) []string
		UpdateCommisionDB(productID int64, commission int) error
	}

	product struct {
		ctx   context.Context
		db    string
		cache string
		es    string
		api   string
	}

	productMock struct {
		db    string
		cache string
		es    string
	}
)

func New(ctx context.Context, src ...string) IProduct {
	var pr IProduct = &productMock{}
	if !mock {
		product := &product{}
		product.ctx = ctx
		pr = product
	}
	return pr
}
