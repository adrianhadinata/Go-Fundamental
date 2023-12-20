package main

import "fmt"

func main() {
	var a int32 = 10
	fmt.Println(a)

	var ab float32 = 0.9
	var b float32 = float32(a) + ab
	fmt.Println(b)

	var c int32 = int32(b)
	fmt.Println(c)
}