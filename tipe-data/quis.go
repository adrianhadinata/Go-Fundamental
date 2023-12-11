package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    )

func quis() {
    var x int
    var y float64

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    x, _ = strconv.Atoi(scanner.Text())
    scanner.Scan()
    y, _ = strconv.ParseFloat(scanner.Text(), 32)

    //Tulis kode disini
    
    var convertedY int = int(y)
    var convertedX float64 = float64(x)
    
    var penguranganFloat = convertedX - y
    
    fmt.Println(x + convertedY)
    fmt.Printf("%.2f" + "\n", penguranganFloat)
    fmt.Println(x * convertedY)
}