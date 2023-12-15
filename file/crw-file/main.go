package main

import "fmt"
import "os"
import "bufio"

var path = `D:\data adrian\Program\Golang\data\`
var filename = "example.txt"
var filePath = path + filename

func main() {
	// var input string
	// scanner := bufio.NewScanner(os.Stdin)

	// fmt.Print("Masukan nama: ")
	// scanner.Scan()
	// input = scanner.Text()

	// write(input)

	// read()

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Dara:", data)
}

// membuat file
func create() {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err){
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		fmt.Println("File", filename, "berhasil dibuat.")
	}
}

// menulis file
func write(input string) {
	// Ini bakal overwrite
	// file, err := os.OpenFile(filePath, os.O_RDWR, 0666)

	// Ini appen
	file, err := os.OpenFile(filePath,os.O_APPEND | os.O_RDWR, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(input + "\n")
	writer.Flush()
}

// Membaca file
func read() {
	file, err := os.OpenFile(filePath,os.O_APPEND | os.O_RDWR, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}