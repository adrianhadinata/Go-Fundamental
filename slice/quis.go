package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
 
func evenNames(slice []string) []string {
    //Tulis kode disini
    var evenChar []string

	for i := 0; i < len(slice); i++ { 

		if len(slice[i]) % 2 == 0 {
			evenChar = append(evenChar, slice[i])
		}
	}

	return evenChar
}
 

func quis() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    x := scanner.Text()
    slice := strings.Split(x, " ")
    names := evenNames(slice)
    fmt.Println(strings.Join(names, " "))
}