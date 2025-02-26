package sorting

import "github.com/e-harsley/scnip_product_catalog_test/models"

type ConversionRateSorter struct {
	BaseSorter
}

func NewConversionRateSorter() *ConversionRateSorter {
	return &ConversionRateSorter{
		BaseSorter{
			name: "Conversion",
			lessFunc: func(p1, p2 *models.Product) bool {
				rate1 := 0.0
				if p1.ViewsCount > 0 {
					rate1 = float64(p1.SalesCount) / float64(p1.ViewsCount)
				}

				rate2 := 0.0
				if p2.ViewsCount > 0 {
					rate2 = float64(p2.SalesCount) / float64(p2.ViewsCount)
				}
				return rate1 > rate2
			},
		},
	}
}
