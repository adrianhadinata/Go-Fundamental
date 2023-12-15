package main

import "fmt"

type mesin struct {
	tipe string
	cc int
}

type kendaraan struct {
	merek string
	tahun int
	model string
	harga int
	mesin
}

func main() {
	var a = kendaraan{
		merek: "Toyota",
		tahun: 2023,
		model: "Camry",
		harga: 500000000,
		mesin: mesin{
			tipe: "premium",
			cc: 2000,
		},
	}

	fmt.Println(a)
	fmt.Println(a.mesin.tipe)
}