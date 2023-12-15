package main 

import "fmt"

func main() {
	var x = 4
	var y = x

	y = y + 1
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("==========================")

	var a = 3
	increase(a)

	fmt.Println(a)

	fmt.Println("==========================")

	var numbers = [4]int {4,3,7,1}
	var anotherNumbers = numbers

	fmt.Println("numbers", numbers)
	fmt.Println("another numbers", anotherNumbers)

	numbers[1] = 9

	fmt.Println("numbers", numbers)
	fmt.Println("another numbers", anotherNumbers)

	fmt.Println("==========================")

	var scores = [4]int {1, 2, 3, 4}
	multiplyByFive(scores)
}

func increase(n int) int {
	n = n + 1
	fmt.Println("n:", n)
	return n
}

func multiplyByFive(n [4]int){
	for i := range n {
		n[i] = n[i] * 5
	}

	fmt.Println("n in function:", n)
}