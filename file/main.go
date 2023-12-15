package main

import "fmt"
import "strconv"

var names []string

func main() {
	var input string
	fmt.Println("names:", names)
	fmt.Print("Masukan jumlah nama: ",)
	fmt.Scanln(&input)

	len, _ := strconv.Atoi(input)

	for i := 0; i < len; i++ {
		fmt.Print("Masukan nama: ")
		fmt.Scanln(&input)

		names = append(names, input)
	}

	fmt.Println("names:", names)
}