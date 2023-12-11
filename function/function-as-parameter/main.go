package main

import "fmt"

func main() {
	fmt.Println("Bilangan Genap ?", isEven(1))
	fmt.Println("Bilangan Ganjil ?", isOdd(1))

	numbers := []int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println("Filter bilangan genap:", filter(numbers, isEven))
	fmt.Println("Filter bilangan ganjil:", filter(numbers, isOdd))
}

func filter(numbers []int, f func(int)bool ) []int {
	var result []int

	for _, number := range numbers {
		if f(number) {
			result = append(result, number)
		}
	}

	return result
}

func isEven(number int) bool {
	return number % 2 == 0
}

func isOdd(number int) bool {
	return number % 2 != 0
}