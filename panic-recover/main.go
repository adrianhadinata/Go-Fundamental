package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("deffered function")
		// if r := recover(); r != nil {
		// 	fmt.Println("Terjadi panic:", r)
		// }
	}()

	fmt.Println("Before panic")
	panic("Panic Occured")

	// var input string
	// fmt.Print("Masukan nama: ")
	// fmt.Scanln(&input)

	// if !isEmpty(input) {
	// 	fmt.Println("Nama anda:", input)
	// }
}

func isEmpty(input string) (empty bool) {
	if input == "" {
		panic("Input tidak boleh kosong")
	}

	empty = false
	return
}