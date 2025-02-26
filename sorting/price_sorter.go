package sorting

import "github.com/e-harsley/scnip_product_catalog_test/models"

type (
	PriceSorter struct {
		BaseSorter
	}
)

func NewPriceSorter() *PriceSorter {
	return &PriceSorter{
		BaseSorter{
			name: "Price",
			lessFunc: func(p1, p2 *models.Product) bool {
				return p1.Price < p2.Price
			},
		},
	}
}
