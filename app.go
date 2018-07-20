package main

import (
	"context"

	"github.com/juliardidev/affiliation/module/seller"
)

func main() {
	seller.AffiliateAll(context.Background(), 10)
}
