package sortingContext

import (
	"github.com/e-harsley/scnip_product_catalog_test/models"
	"github.com/e-harsley/scnip_product_catalog_test/sorting"
)

type SortingContext struct {
	strategy sorting.Sorter
}

func NewSortingContext(strategy sorting.Sorter) *SortingContext {
	return &SortingContext{
		strategy: strategy,
	}
}

func (c *SortingContext) SetStrategy(strategy sorting.Sorter) {
	c.strategy = strategy
}

func (c *SortingContext) Sort(products []*models.Product) {
	c.strategy.Sort(products)
}

type SortingRegistry struct {
	strategies map[string]sorting.Sorter
}

func NewSortingRegistry() *SortingRegistry {
	return &SortingRegistry{
		strategies: make(map[string]sorting.Sorter),
	}
}

func (r *SortingRegistry) Register(strategy sorting.Sorter) {
	r.strategies[strategy.SortingBy()] = strategy
}

func (r *SortingRegistry) GetStrategy(name string) (sorting.Sorter, bool) {
	strategy, exists := r.strategies[name]
	return strategy, exists
}

func (r *SortingRegistry) GetAllStrategies() []sorting.Sorter {
	strategies := make([]sorting.Sorter, 0, len(r.strategies))
	for _, strategy := range r.strategies {
		strategies = append(strategies, strategy)
	}
	return strategies
}
