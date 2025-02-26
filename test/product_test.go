package test

import (
	"bytes"
	"github.com/e-harsley/scnip_product_catalog_test/models"
	"github.com/e-harsley/scnip_product_catalog_test/sorting"
	"github.com/e-harsley/scnip_product_catalog_test/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var testProductsJSON = `
[
    {
        "id": 1,
        "name": "Alabaster Table",
        "price": 12.99,
        "created_at": "2019-01-04T00:00:00Z",
        "sale_count": 32,
        "view_count": 730
    },
    {
        "id": 2,
        "name": "Zebra Table",
        "price": 44.49,
        "created_at": "2012-01-04T00:00:00Z",
        "sale_count": 301,
        "view_count": 3279
    },
    {
        "id": 3,
        "name": "Coffee Table",
        "price": 10.00,
        "created_at": "2014-05-28T00:00:00Z",
        "sale_count": 1048,
        "view_count": 20123
    }
]`

func createTestProductsFromJSON() ([]*models.Product, error) {
	products := []*models.Product{}
	reader := bytes.NewReader([]byte(testProductsJSON))
	err := utils.BindDataOperationStruct(reader, &products)
	return products, err
}

func createTestJSONFile(t *testing.T) string {
	tmpFile, err := os.CreateTemp("", "products-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	_, err = tmpFile.Write([]byte(testProductsJSON))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	err = tmpFile.Close()
	if err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tmpFile.Name()
}

func TestSortersImplementInterface(t *testing.T) {
	priceSorter := sorting.NewPriceSorter()
	var sorter sorting.Sorter = priceSorter
	assert.NotNil(t, sorter)
	assert.Equal(t, "Price", sorter.SortingBy())

	conversionSorter := sorting.NewConversionRateSorter()
	sorter = conversionSorter
	assert.NotNil(t, sorter)
	assert.Equal(t, "Conversion", sorter.SortingBy())

	dateSorter := sorting.NewDateSorter()
	sorter = dateSorter
	assert.NotNil(t, sorter)
	assert.Equal(t, "Date", sorter.SortingBy())
}

func TestJSONParsing(t *testing.T) {
	products, err := createTestProductsFromJSON()
	assert.NoError(t, err)
	assert.Len(t, products, 3)

	assert.Equal(t, 1, products[0].ID)
	assert.Equal(t, "Alabaster Table", products[0].Name)
	assert.Equal(t, 12.99, products[0].Price)
	assert.Equal(t, 32, products[0].SalesCount)
	assert.Equal(t, 730, products[0].ViewsCount)

	expectedTime, _ := time.Parse(time.RFC3339, "2019-01-04T00:00:00Z")
	assert.Equal(t, expectedTime, products[0].CreatedAt)
}

func TestLoadFromFile(t *testing.T) {
	filename := createTestJSONFile(t)
	defer os.Remove(filename)

	file, err := os.Open(filename)
	assert.NoError(t, err)
	defer file.Close()

	products := []*models.Product{}
	err = utils.BindDataOperationStruct(file, &products)
	assert.NoError(t, err)
	assert.Len(t, products, 3)
}

func TestPriceSorting(t *testing.T) {
	products, err := createTestProductsFromJSON()
	assert.NoError(t, err)

	priceSorter := sorting.NewPriceSorter()
	priceSorter.Sort(products)

	assert.Equal(t, 3, products[0].ID, "First product should be Coffee Table (ID 3)")
	assert.Equal(t, 1, products[1].ID, "Second product should be Alabaster Table (ID 1)")
	assert.Equal(t, 2, products[2].ID, "Third product should be Zebra Table (ID 2)")
}

func TestConversionRateSorting(t *testing.T) {
	products, err := createTestProductsFromJSON()
	assert.NoError(t, err)

	conversionSorter := sorting.NewConversionRateSorter()
	conversionSorter.Sort(products)

	rates := make(map[int]float64)
	for _, p := range products {
		if p.ViewsCount > 0 {
			rates[p.ID] = float64(p.SalesCount) / float64(p.ViewsCount)
		} else {
			rates[p.ID] = 0
		}
	}

	for i := 0; i < len(products)-1; i++ {
		assert.GreaterOrEqual(t,
			rates[products[i].ID],
			rates[products[i+1].ID],
			"Products should be sorted by descending conversion rate")
	}
}

func TestDateSorting(t *testing.T) {
	products, err := createTestProductsFromJSON()
	assert.NoError(t, err)

	dateSorter := sorting.NewDateSorter()
	dateSorter.Sort(products)

	assert.Equal(t, 1, products[0].ID, "First product should be Alabaster Table (ID 1)")
	assert.Equal(t, 3, products[1].ID, "Second product should be Coffee Table (ID 3)")
	assert.Equal(t, 2, products[2].ID, "Third product should be Zebra Table (ID 2)")
}
