package main

import "fmt"

func main() {
	var a int = 10
	a = a - 3

	fmt.Println(a)

	var b int = 5
	var c int

	c = b
	fmt.Println(c)
	c += b
	fmt.Println("Tambah =", c)
	c -= b
	fmt.Println("Kurang =", c)
	c *= b
	fmt.Println("Kali =", c)
	c /= b
	fmt.Println("Bagi =", c)
	c %= b
	fmt.Println("Modulus =", c)
}