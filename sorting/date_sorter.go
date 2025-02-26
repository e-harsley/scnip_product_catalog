package sorting

import "github.com/e-harsley/scnip_product_catalog_test/models"

type DateSorter struct {
	BaseSorter
}

func NewDateSorter() *DateSorter {
	return &DateSorter{
		BaseSorter{
			name: "Date",
			lessFunc: func(p1, p2 *models.Product) bool {
				return p1.CreatedAt.After(p2.CreatedAt)
			},
		},
	}
}
