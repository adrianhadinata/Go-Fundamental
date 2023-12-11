package main

import "fmt"

func main() {
	numbers := []int {2, 4, 1, 10, 9}

	kecil, besar := minMax(numbers)

	fmt.Println("Kecil:", kecil)
	fmt.Println("Besar:", besar)
}

func minMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return min, max
}