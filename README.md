# Go Price Calculator

This project is a price calculator implemented in Go, developed as part of the "Go - The Complete Guide" course by Maximilian Schwarzmuller on Udemy (Section 9).

## Project Overview

The Go Price Calculator is an application that calculates prices with different tax rates. It demonstrates various Go programming concepts, including interfaces, file I/O, and command-line interaction.

## Files

- `price_calculator.go`: Main application entry point
- `prices/prices.go`: Core price calculation logic
- `filemanager/filemanager.go`: File I/O operations
- `iomanager/iomanager.go`: I/O manager interface
- `cmdmanager/cmdmanager.go`: Command-line interface manager
- `conversion/conversion.go`: Utility functions for data conversion

## Features

- Calculate prices with various tax rates
- Flexible input options:
  - Read input prices from a file
  - Accept input prices via command-line
- Flexible output options:
  - Write results to JSON files
  - Display results in the console
- Modular design using interfaces for I/O management

## Usage

The application can be run in two modes:

1. Command-line mode:
   ```
   go run price_calculator.go
   ```
   Follow the on-screen prompts to enter prices and view the calculated results in the console.

2. File I/O mode:
   Uncomment and modify the relevant lines in `price_calculator.go` to use file input and output:
   ```go
   fm, err := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
   if err != nil {
       fmt.Println(err)
       return
   }
   job := prices.NewTaxIncludedPriceJob(fm, taxRate)
   ```
   Ensure you have an input file (e.g., `prices.txt`) with prices listed one per line. Results will be written to JSON files.

   Make sure to comment out or remove the `cmdmanager` related lines when using file I/O mode.

## Course Information

This project was completed as part of the "Go - The Complete Guide" course by Maximilian Schwarzmuller on Udemy. You can find the course at:

[https://www.udemy.com/course/go-the-complete-guide/](https://www.udemy.com/course/go-the-complete-guide/)

## License

This project is for educational purposes and follows the licensing terms of the original course material.
