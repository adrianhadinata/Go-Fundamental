package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
 
func quis() {
   // Tulis kode disini

   scanner := bufio.NewScanner(os.Stdin)
   scanner.Scan()

   capacity, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

   var arr = make([]int, capacity)

    scanner.Scan()
    inputNumber := strings.Split(scanner.Text(), " ")

    for i := 0; i < len(inputNumber); i++ {
        arr[i], _ = strconv.Atoi(inputNumber[i])
    }

    for _, inputNumber := range arr {
        if inputNumber % 2 == 0 && inputNumber != 0 {
            fmt.Println(inputNumber)
        }
    }
}