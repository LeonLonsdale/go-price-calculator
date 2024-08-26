package prices

import (
	"bufio"
	"fmt"
	"os"

	"github.com/LeonLonsdale/go-price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	// Open the "prices.txt" file
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("An error occurred while opening the file:")
		fmt.Println(err)
		return
	}

	// Ensure the file is closed when the function exits
	defer file.Close()

	// Create a new Scanner to read the file
	scanner := bufio.NewScanner(file)

	var lines []string
	// Read the file line by line
	for scanner.Scan() {
		// Append each line to the lines slice
		lines = append(lines, scanner.Text())
	}

	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("An error occurred while reading the file:")
		fmt.Println(err)
		return
	}

	prices, error := conversion.StringsToFloat(lines)
	if error != nil {
		fmt.Println("An error occurred while converting price string to float:")
		fmt.Println(error)
		file.Close()
		return
	}

	// receiver arg must be pointer so that this is persisted to the job and not a copy.
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
}
