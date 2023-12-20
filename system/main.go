package main

import (
	"fmt"
	"runtime"
)

func main() {
	sistemOperasi := runtime.GOOS
	fmt.Println(sistemOperasi)
}