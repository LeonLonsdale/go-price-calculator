package main

import "fmt"

type floatSliceMap map[float64][]float64

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(floatSliceMap, 4)

	for _, tr := range taxRates {
		newPrices := make([]float64, len(prices))
		for _, p := range prices {
			price := p + (p * tr)
			// could also be: price := p * (1 + tr)
			newPrices = append(newPrices, price)
		}
		result[tr] = newPrices
	}

	fmt.Println(result)
}
