package main

import "fmt"

type Currency struct {
	CurrencyCode string
	CurrencyName string
}

func main() {
	currConvMap := map[Currency]float64{
		{"USD", "United States Dollar"}:  1.1318,
		{"JPY", "Japanese Yen"}:         121.05,
		{"GBP", "British Pound"}:        0.90630,
		{"CNY", "Chinese Yuan"}:         7.9944,
		{"SGD", "Singapore Dollar"}:     1.5743,
		{"MYR", "Malaysian Ringgit"}:    4.8390,
	}

	for curr, rate := range currConvMap {
		fmt.Printf("%s (%s) = %.4f\n", curr.CurrencyCode, curr.CurrencyName, rate)
	}
}