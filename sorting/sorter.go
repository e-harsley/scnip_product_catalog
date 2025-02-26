package sorting

import "github.com/e-harsley/scnip_product_catalog_test/models"

type Sorter interface {
	Sort(products []*models.Product)
	SortingBy() string
}
