package seller

import (
	"context"
	"log"

	"github.com/juliardidev/affiliation/module/seller/internal/product"
)

func AffiliateAll(ctx context.Context, limit int64) {
	product := product.New(ctx, "postgres", "redis", "elasticsearch")
	log.Println(product.FetchAllDB(limit))
}

func AffiliateProduct(productID int64) {

}

func AffiliateProducts(productIDs []int64) {

}
