package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello World")
	}()

	kata := "Halo Dunia"

	halo := func(word string) {
		fmt.Println(kata)
	}

	halo(kata)
}