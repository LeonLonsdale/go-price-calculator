package main

import (
	"fmt"

	"github.com/LeonLonsdale/go-price-calculator/filemanager"
	"github.com/LeonLonsdale/go-price-calculator/prices"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm, err := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		if err != nil {
			fmt.Println(err)
			return
		}
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		job.Process()
	}
}
