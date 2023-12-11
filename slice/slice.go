package main

import "fmt"

func main() {
	// Ini array
	ages := [4]int {20, 21, 22, 23}

	// Buat slice
	slicedAges := ages[0:3]

	fmt.Println(ages)
	fmt.Println(slicedAges)

	// Pass by reference

	slicedAges = append(slicedAges, 37)

	fmt.Println(ages)
	fmt.Println(slicedAges)
}
