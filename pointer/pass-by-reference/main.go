package main

import "fmt"

func main() {
	var x int = 4

	// variable khusus untuk menyimpan pointer
	var y *int = &x

	fmt.Println("x", x)
	fmt.Println("alamat dari x", &x)
	fmt.Println("y", y)
	fmt.Println("alamat dari y", &y)

	fmt.Println("Nilai referensi dari pointer y:", *y)

	*y = *y + 1

	fmt.Println("x:", x)
	fmt.Println("==========================")

	var a int = 3
	increase(&a)
	fmt.Println("a:", a)

	fmt.Println("==========================")

	var numbers = []int{4, 3 , 7, 11} 
	var anotherNumbers = numbers 

	fmt.Println("numbers:", numbers)
	fmt.Println("Another numbers:", anotherNumbers)

	numbers[1] = 9

	fmt.Println("numbers:", numbers)
	fmt.Println("Another numbers:", anotherNumbers)

	var scores = []int {7, 8 ,6 ,9}
	multiplyByFive(scores)
	fmt.Println("scores:", scores)
}

func increase(n *int) {
	*n = *n + 1
	fmt.Println("n di function increase:", *n)
}

func multiplyByFive(numbers []int) {
	for i := range numbers {
		numbers[i] = numbers[i] * 5
	}

	fmt.Println("Numbers di function: ", numbers)
}