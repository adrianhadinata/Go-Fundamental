package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println("==========================")
	i := 6

	for i <= 10 {
		fmt.Println("Angka", i)
		i++ 
	}

	fmt.Println("==========================")

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			// Print untuk menyamping
			fmt.Print(j, " ")
		}
		// Untuk pindah baris
		fmt.Println()
	}

	fmt.Println("==========================")

	
}