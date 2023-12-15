package main

import "fmt"

func main() {
	defer fmt.Println("Start")
	fmt.Println("Process")
	fmt.Println("Done")
}