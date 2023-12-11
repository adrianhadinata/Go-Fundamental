package main

import "fmt"

func main() {
	fmt.Println("Hasil penjumlahan:", add(3, 4))
	fmt.Println("Hasil perkalian:", multiply(3, 4))
	fmt.Println("Hasil penggabungan:", add(3, multiply(3, 4)))
}

func add(a int, b int) int {
	hasil := a + b
	return hasil
}

func multiply(a, b int) int {
	hasil := a * b
	return hasil
}