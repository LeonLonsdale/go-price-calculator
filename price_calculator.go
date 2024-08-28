package main

import (
	"fmt"

	"github.com/LeonLonsdale/go-price-calculator/filemanager"
	"github.com/LeonLonsdale/go-price-calculator/prices"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		fm, err := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		// cmd, err := cmdmanager.New()
		if err != nil {
			fmt.Println(err)
			return
		}

		job := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go job.Process(doneChans[i], errorChans[i])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
