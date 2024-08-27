package prices

import (
	"fmt"

	"github.com/LeonLonsdale/go-price-calculator/conversion"
	"github.com/LeonLonsdale/go-price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"prices"`
	TaxIncludedPrices map[string]string   `json:"prices_with_tax"`
	IOManager         iomanager.IOManager `json:"-"`
}

func NewTaxIncludedPriceJob(iomanager iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   iomanager,
	}
}

// receiver arg must be pointer so that this is persisted to the job and not a copy.
func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, error := conversion.StringsToFloat(lines)
	if error != nil {
		fmt.Println("An error occurred while converting price string to float:")
		fmt.Println(error)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
}
