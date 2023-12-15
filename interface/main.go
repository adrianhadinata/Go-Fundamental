package main

import "fmt"

type Rectangle struct {
	width float64
	length float64
}

func (r Rectangle) luas() float64 {
	return r.width * r.length
}

func (r Rectangle) keliling() float64 {
	return 2 * (r.width * r.length)
}

type Kotak struct {
	side float64
}

func (k Kotak) luas() float64 {
	return k.side * k.side
}

func (k Kotak) keliling() float64 {
	return k.side * 4
}

// struct dapat dikatakan bagian dari suatu interface apabila method yang dimiliki sama dengan interface tersebut

type Shape interface {
	luas() float64
	keliling() float64
}

func main() {

	// Upcasting

	var shape1 Shape = Kotak{side: 5}
	var shape2 Shape = Rectangle{width: 4, length: 5}
	var shape3 Shape = Rectangle{width: 7, length: 9}

	fmt.Println("shape 1:", shape1)
	fmt.Println("shape 2:", shape2)
	fmt.Println("shape 3:", shape3)

	shapes := []Shape{shape1, shape2, shape3}

	for _, shape := range shapes{
		fmt.Printf("%#v memiliki luas  %.1f dan keliling %.1f \n", shape, shape.luas(), shape.keliling())
	}

	fmt.Println("=================================================")

	var anything interface {}

	anything = "Adrian"
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = 20
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = false
	fmt.Printf("anything %T: %v \n", anything, anything)
}