package main

import "fmt"

func main() {
	fmt.Println("Tipe Data")
	fmt.Println("Numeric")
	var positiveNumber uint8 = 90
	var positiveNumberInt int = 9
	fmt.Println("Bilangan positif:", positiveNumber, positiveNumberInt)
	var negativeNumber int = -9
	fmt.Println("Bilangan negative:", negativeNumber)
	var decimalNumber = 3.90
	fmt.Printf("Bilangan desimal %f lalu batasin nol dengan cara %.2f \n", decimalNumber, decimalNumber)
}