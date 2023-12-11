package main

import "fmt"

func main() {
	result := "lulus";

	if result == "lulus" {
		fmt.Println("Selamat anda", result)
	} else {
		fmt.Println("maaf anda", result)
	}

	GPA := 2.40

	if GPA >= 3.50 && GPA <= 4.00 {
		result = "CUMLAUDE"
	} else if GPA >= 3.00 && GPA <= 3.49 {
		result = "MEMUASKAN"
	} else if GPA >= 2.99 && GPA <= 2.50 {
		result = "CUKUP MEMUASKAN"
	} else {
		result = "PERLU PERHATIAN KHUSUS"
	}

	fmt.Printf("GPA anda %.2f, Selamat anda lulus dengan predikat %s \n", GPA, result)
	fmt.Println("======================================================")

	poin := 7

	switch poin {
		case 8: {
			fmt.Println("Bagus")
		}
		case 7: {
			fmt.Println("Cukup")
		}
		default: {
			fmt.Println("Kurang")
		}
	}
}