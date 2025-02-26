package main

import (
	"fmt"
	sortingContext "github.com/e-harsley/scnip_product_catalog_test/context"
	"github.com/e-harsley/scnip_product_catalog_test/models"
	"github.com/e-harsley/scnip_product_catalog_test/sorting"
	"github.com/e-harsley/scnip_product_catalog_test/utils"
	"os"
)

func RegisterSortingStrategy() *sortingContext.SortingRegistry {
	registry := sortingContext.NewSortingRegistry()
	registry.Register(sorting.NewPriceSorter())
	registry.Register(sorting.NewConversionRateSorter())
	registry.Register(sorting.NewDateSorter())
	return registry
}

func main() {

	file, err := os.Open("products.json")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	products := []*models.Product{}

	err = utils.BindDataOperationStruct(file, &products)

	if err != nil {
		fmt.Println("Failed to bind file data to products", err)
		return
	}

	registry := RegisterSortingStrategy()

	priceSorter, ok := registry.GetStrategy("Price")

	if !ok {
		fmt.Println("Failed to get price sorter")
		return
	}

	context := sortingContext.NewSortingContext(priceSorter)
	context.Sort(products)
	models.PrintProducts(products)

	dateSorter, ok := registry.GetStrategy("Date")

	if !ok {
		fmt.Println("Failed to get date sorter")
		return
	}

	context = sortingContext.NewSortingContext(dateSorter)
	context.Sort(products)
	models.PrintProducts(products)

	conversionSorter, ok := registry.GetStrategy("Conversion")

	if !ok {
		fmt.Println("Failed to get date sorter")
		return
	}

	context = sortingContext.NewSortingContext(conversionSorter)
	context.Sort(products)
	models.PrintProducts(products)
}
