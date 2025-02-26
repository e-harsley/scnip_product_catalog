package sorting

import (
	"github.com/e-harsley/scnip_product_catalog_test/models"
	"sort"
)

type (
	BaseSorter struct {
		name     string
		lessFunc func(p1, p2 *models.Product) bool
	}
)

func (s *BaseSorter) SortingBy() string {
	return s.name
}

func (s *BaseSorter) Sort(products []*models.Product) {
	sort.Slice(products, func(i, j int) bool {
		return s.lessFunc(products[i], products[j])
	})
}
