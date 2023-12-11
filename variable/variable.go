package main

import "fmt"

func main() {
	var firstName string = "Adrian"
	var lastName string = "Hadinata Hadi Darsono"

	fmt.Println(firstName, lastName)
	fmt.Printf("Mencetak ke terminal dengan format tertentu! (mencetak dalam 1 row) \n")
	fmt.Printf("Hallo, %s %s", firstName, lastName)

	var (
		age int
		address string
	)

	age = 24
	address = "Semarang"

	fmt.Println("Saya berumur", age, "dan tinggal di", address)
	
	var (
		bootcampName,  bootcampAddress = 
		"Enigma", "Jakarta"
	)

	fmt.Println("Saya sedang mengikuti bootcamp di", bootcampName, bootcampAddress)

	// deklarasi tipe variable mengikuti nilai di dalamnya
	month := "Desember"
	year := "2023"

	fmt.Println("Saya mulai mendaftar pada", month, year)
}