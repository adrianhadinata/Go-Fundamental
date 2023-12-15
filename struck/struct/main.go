package main

import "fmt"

type kendaraan struct {
	merek string
	tahun int
	model string
	harga int
}

func main() {
	var a kendaraan
	a.merek = "Toyota"
	a.tahun = 2023
	a.model = "Innova"
	a.harga = 300000000

	fmt.Println("a:", a)

	var b = kendaraan {"Wuling", 2022, "Confero", 250000000}
	fmt.Println("b:", b)

	var c = kendaraan {merek:"Honda", model: "Brio", harga: 250000000, tahun: 2022}
	fmt.Println("c:", c)

	fmt.Println("=========================")

	var x = kendaraan {merek:"Toyota", model: "Agya", harga: 100000000, tahun: 2019}
	fmt.Println("x:", x)
	fmt.Printf("alamat x: %p\n", &x) // hanya bisa di print f

	var y = x
	fmt.Println("y:", y)
	fmt.Printf("alamat y: %p\n", &y)

	y.model = "Kijang"

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("=========================")

	var k = new(kendaraan)
	fmt.Printf("isi k: %p\n", k)





}