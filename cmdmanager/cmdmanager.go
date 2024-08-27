package cmdmanager

import "fmt"

type Cmdmanager struct{}

func New() (Cmdmanager, error) {
	return Cmdmanager{}, nil
}

func (cmd Cmdmanager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm each price with ENTER")

	var prices []string

	for {
		var price string
		var option string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		prices = append(prices, price)
		fmt.Print("Add another? [1] Yes, [2] No: ")
		fmt.Scan(&option)
		if option == "2" {
			break
		}
	}
	return prices, nil
}

func (cmd Cmdmanager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
