package main

import "fmt"

func main() {
	var fruits = [4]string {
		"Apel", "Pisang", "Anggur", "Semangka"}

	fmt.Println(fruits)
	fmt.Println(fruits[1])

	fruits[2] = "Jeruk"

	fmt.Println(fruits)

	fmt.Println("======================================================")

	var days = [...]string {"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	fmt.Println(days)
	fmt.Println("Jumlah item array :", len(days))

	fmt.Println("======================================================")

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("======================================================")

	var matrix = [2][3]int {
		{3,2,3}, 
		{3,4,5},
	}

	fmt.Println(matrix)
}



// for i := 0; i < capacity; i++ {
//     scanner.Scan()
//     inputNumber, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

//     arr = append(arr, inputNumber)
// }

//    var capacity int
//    fmt.Scanln(&capacity)

//    var arr = make([]int, capacity)
   
//    for i := 0; i < capacity; i++ {
//        fmt.Scanln(&arr[i])
//    }
   
//    for i := 0; i < len(arr); i++ {
//        if arr[i] % 2 == 0 {
//            fmt.Println(arr[i])
//        }
//    }