package main

import (
    "bufio"
    "fmt"
    "os"
)

func quis() {
    //Tulis kode disini
    var nama string
    var alamat string
    var email string
    
    scanner := bufio.NewScanner(os.Stdin)
    
    scanner.Scan()
    nama = scanner.Text()
    scanner.Scan()
    alamat = scanner.Text()
    scanner.Scan()
    email = scanner.Text()
    
    fmt.Printf("Halo, saya %s. Saya tinggal di %s. Alamat email saya adalah %s", nama, alamat, email)
}