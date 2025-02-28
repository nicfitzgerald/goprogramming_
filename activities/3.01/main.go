package main

import "fmt"

func main() {
	total := 0.0
	total += salesTaxCalc(0.99, .075)
	fmt.Println(total)
	total += salesTaxCalc(2.75, .015)
	fmt.Println(total)
	total += salesTaxCalc(0.87, .02)
	fmt.Println(total)

	fmt.Println("\nSales tax total: ", total)
}

func salesTaxCalc(total float64, tax float64) float64 {
	return total * tax
}
