package models

import (
	"fmt"
	"time"
)

type (
	Product struct {
		ID         int       `json:"id"`
		Name       string    `json:"name"`
		Price      float64   `json:"price"`
		CreatedAt  time.Time `json:"created_at"`
		SalesCount int       `json:"sale_count"`
		ViewsCount int       `json:"view_count"`
	}
)

func PrintProducts(products []*Product) {
	fmt.Println("ID\tName\t\t\tPrice\tCreated\t\tSales\tViews\tConversion Rate")
	fmt.Println("----------------------------------------------------------------------------------")
	for _, p := range products {
		convRate := 0.0
		if p.ViewsCount > 0 {
			convRate = float64(p.SalesCount) / float64(p.ViewsCount)
		}
		fmt.Printf("%d\t%-20s\t$%.2f\t%s\t%d\t%d\t%.4f\n",
			p.ID, p.Name, p.Price, p.CreatedAt.Format("2006-01-02"), p.SalesCount, p.ViewsCount, convRate)
	}
	fmt.Println()
}
