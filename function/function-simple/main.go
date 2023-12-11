package main

import "fmt"

func main() {
	helloWorld()
	// greet("Adrian")
	greet("Adrian", 2)
	fmt.Println(multiplier(2, 2))
}

func helloWorld() {
	fmt.Println("Hello World!")
}

func greet(name string, tanggal int) {
	fmt.Println("Hello", name, "Tanggal", tanggal)
}

func multiplier(angka1 int, angka2 int) int {
	hasil := angka1 * angka2
	return hasil
}