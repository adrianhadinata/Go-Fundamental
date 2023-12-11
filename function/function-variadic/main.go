package main

import "fmt"

func main() {
	fmt.Println("Hasil penjumlahan:", sum(2,3,4,5,6))
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}