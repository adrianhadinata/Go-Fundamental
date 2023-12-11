package main

import "fmt"

func quis() {
    // Tulis kode disini
    var looping int
    fmt.Scanln(&looping)
    
    increment := looping
    
    for increment >= looping {
        fmt.Println(looping, " I will become a go developer")
        looping--
        
        if looping == 0 {
            break
        }
    }
}