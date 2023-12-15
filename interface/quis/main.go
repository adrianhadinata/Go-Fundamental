package main

import "fmt"
import "bufio"
import "os"
import "strconv"

type BangunDatar interface {
    Luas() int
}

//kode struct disini
type Segitiga struct {
    alas int
	tinggi int
}

//kode struct method disini
func (s Segitiga) Luas() int {
	return s.alas * s.tinggi / 2
}

func getLuas(bangunDatar BangunDatar) {
    fmt.Printf("Luas bangun datar = %d \n", bangunDatar.Luas())
}

func main() {
   var segitiga1 Segitiga
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    segitiga1.alas, _ = strconv.Atoi(scanner.Text())
    scanner.Scan()
    segitiga1.tinggi, _ = strconv.Atoi(scanner.Text())
    getLuas(&segitiga1)

}