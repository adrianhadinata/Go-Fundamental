package main

import "fmt"

type Patient struct {
	Name string
	Age int
	Celsius
}

type Celsius float64 

func (c Celsius) toFrenheit() float64 {
	return float64(c * 9 / 5 +32)
}

func main() {
	var temperature Celsius = 20.0
	fmt.Println("temperature", temperature)
	fmt.Println("suhu dalam farenheit", temperature.toFrenheit())

	fmt.Println("================================================")

	newPatient := Patient{Name: "Adrian", Age: 24, Celsius: 39.0,}
	fmt.Println(newPatient)
}