package product

import (
	"context"
	"testing"
)

func TestFetch(t *testing.T) {
	product := New(context.Background())

	t.Error(product.FetchAllDB(10))
}
