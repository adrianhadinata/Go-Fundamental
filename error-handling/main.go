package main

import "fmt"
import "errors"

func main() {

	// contoh output panic
	// result := divide(5, 0)

	result, error := divide(5, 1)

	if error != nil {
		fmt.Println("Error:", error)
		return
	}
	
	fmt.Println(result)
}

func divide (a,b int) (int, error) {
	if b == 0 {
		return 0, errors.New("pembagian dengan angka 0")
	}

	return a / b, nil
}