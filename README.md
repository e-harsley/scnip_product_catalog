# Product Catalog Sorting Solution

This repository contains a flexible and extensible product catalog sorting system implemented in Go. The solution allows different teams to add their own sorting strategies without modifying existing code.

## Problem Statement

Our product manager wants to A/B test different product sorting methods to determine which one drives the most sales. Different parts of the website might benefit from different sorting methods, so the system needs to be highly extensible.

## Solution

The implementation uses several design patterns to create a clean, maintainable solution:

1. **Strategy Pattern** - Encapsulates different sorting algorithms in separate classes.
2. **Context Pattern** - Manages and switches between different sorting strategies at runtime.
3. **Registry Pattern** - Allows for dynamic registration of new sorting strategies.
4. **Factory Pattern** - Creates and configures sorting strategies.

## Project Structure

```
scnip_product_catalog_test/
├── context/
│   └── sorting_context.go    # Context and registry implementation
├── models/
│   └── product.go            # Product data structure
├── sorting/
│   ├── base.go               # Base sorter implementation
│   ├── conversion_sorter.go  # Conversion rate sorter
│   ├── date_sorter.go        # Date sorter
│   ├── price_sorter.go       # Price sorter
│   └── sorter.go             # Sorter interface
├── test/
│   └── product_test.go       # Test cases
├── utils/
│   └── utils.go              # Utility functions
├── go.mod                    # Go module definition
├── main.go                   # Example usage
└── products.json             # Sample product data
```

## Features

- **Price Sorting**: Sort products by price (lowest to highest).
- **Conversion Rate Sorting**: Sort by sales-to-views ratio (highest to lowest).
- **Date Sorting**: Sort by creation date (newest first).
- **Extensible Architecture**: New sorting strategies can be added without modifying existing code.

## How to Use

### Basic Usage

```go
package main

import (
    sortingContext "github.com/e-harsley/scnip_product_catalog_test/context"
    "github.com/e-harsley/scnip_product_catalog_test/models"
    "github.com/e-harsley/scnip_product_catalog_test/sorting"
)

func main() {
    // Get products from database or JSON
    products := []*models.Product{...}
    
    // Create a registry and register standard sorting strategies
    registry := sortingContext.NewSortingRegistry()
    registry.Register(sorting.NewPriceSorter())
    registry.Register(sorting.NewConversionRateSorter())
    registry.Register(sorting.NewDateSorter())
    
    // Get a specific sorter and create a context
    priceSorter, _ := registry.GetStrategy("Price")
    context := sortingContext.NewSortingContext(priceSorter)
    
    // Sort products by price
    context.Sort(products)
    
    // Switch to conversion rate sorting
    conversionSorter, _ := registry.GetStrategy("Conversion")
    context.SetStrategy(conversionSorter)
    context.Sort(products)
}
```

## Adding a New Sorting Strategy

Teams can add their own sorting strategies without modifying the existing codebase:

```go
// Create a new sorter implementation
type PopularitySorter struct {
    sorting.BaseSorter
}

func NewPopularitySorter() *PopularitySorter {
    return &PopularitySorter{
        BaseSorter: sorting.BaseSorter{
            Name: "Popularity",
            LessFunc: func(p1, p2 *models.Product) bool {
                return p1.SalesCount > p2.SalesCount // Highest sales first
            },
        },
    }
}

// Register the new sorter with the registry
registry.Register(NewPopularitySorter())

// Use the new sorter
popularitySorter, _ := registry.GetStrategy("Popularity")
context.SetStrategy(popularitySorter)
context.Sort(products)
```

## Running the Code

1. Clone the repository:
    ```sh
    git clone https://github.com/e-harsley/scnip_product_catalog_test.git
    cd scnip_product_catalog_test
    ```
2. Run the example:
    ```sh
    go run main.go
    ```

## Testing

To run the tests:

```sh
go test -v ./test/...
```

## Design Decisions

### Why Strategy Pattern?

- Encapsulates different sorting algorithms.
- Allows switching strategies at runtime.
- Makes adding new strategies simple.

### Why Registry Pattern?

- Provides a central location for managing strategies.
- Allows dynamic registration of new strategies.
- Decouples strategy creation from strategy usage.

### Why Base Sorter?

- Reduces code duplication.
- Simplifies the creation of new sorters.
- Ensures consistent behavior across sorters.

## Technical Requirements

- Go 1.13 or higher.
- [github.com/stretchr/testify](https://github.com/stretchr/testify) (for tests).