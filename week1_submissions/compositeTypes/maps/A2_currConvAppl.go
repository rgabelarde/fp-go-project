package main

import "fmt"

type Currency struct {
	CurrencyCode string
	CurrencyName string
}

func main() {
	var convertFrom, convertTo string
	var amountToConvert float64

	currConvMap := map[Currency]float64{
		{"USD", "United States Dollar"}:  1.1318,
		{"JPY", "Japanese Yen"}:         121.05,
		{"GBP", "British Pound"}:        0.90630,
		{"CNY", "Chinese Yuan"}:         7.9944,
		{"SGD", "Singapore Dollar"}:     1.5743,
		{"MYR", "Malaysian Ringgit"}:    4.8390,
	}

	fmt.Println("Currency Conversion Rates")
	for curr, rate := range currConvMap {
		fmt.Printf("%s (%s) = %.4f\n", curr.CurrencyCode, curr.CurrencyName, rate)
	}

	fmt.Print("Enter the currency code to convert from: ")
	fmt.Scanf("%s\n", &convertFrom)

	var fromCurrency Currency
	validFrom := false
	for curr := range currConvMap {
		if curr.CurrencyCode == convertFrom {
			fromCurrency = curr
			validFrom = true
			break
		}
	}
	if !validFrom {
		fmt.Println("Invalid currency code")
		return
	}

	fmt.Print("Enter the amount to convert: ")
	fmt.Scanf("%f\n", &amountToConvert)
	if amountToConvert <= 0 {
		fmt.Println("Invalid amount")
		return
	}

	fmt.Print("Enter the currency code to convert to: ")
	fmt.Scanf("%s\n", &convertTo)

	var toCurrency Currency
	validTo := false
	for curr := range currConvMap {
		if curr.CurrencyCode == convertTo {
			toCurrency = curr
			validTo = true
			break
		}
	}
	if !validTo {
		fmt.Println("Invalid currency code")
		return
	}

	convertedAmount := (amountToConvert / currConvMap[fromCurrency]) * currConvMap[toCurrency]
	fmt.Printf("%.2f %s (%s) = %.2f %s (%s)\n", amountToConvert, fromCurrency.CurrencyCode, fromCurrency.CurrencyName, convertedAmount, toCurrency.CurrencyCode, toCurrency.CurrencyName)
}
