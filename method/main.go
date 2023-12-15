package main

import "fmt"

/* type Triangel struct {
	alas float64
	tinggi float64
}

// value receiver
func (t Triangel) luas() {
	fmt.Println("Luas segitiga:", 0.5 * t.alas * t.tinggi)
}

// pointer receiver
func (t *Triangel) increase() {
	t.alas += 1
	t.tinggi += 1
}

func main() {
	instance := Triangel{10, 12}
	instance.luas()

	fmt.Println("Sebelum di tingkatkan:", instance)
	instance.increase()
	fmt.Println("Sesudah di tingkatkan:", instance)
} */

type Counter struct {
	Value int
}

func (c Counter) Increment() {
	c.Value++
}

func main() {
	counter := &Counter{Value: 3}
	counter.Increment()
	fmt.Println("Value setelah di Increment:", counter.Value)
}

