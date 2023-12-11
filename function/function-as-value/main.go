package main

import "fmt"

func main() {
	var sum func (int, int) int = add

	fmt.Println("Hasil penjumlahan:", sum(2,3))
}

func add (a int, b int) int {
	hasil := a + b
	return hasil
}