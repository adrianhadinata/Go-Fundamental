package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
	var firstName string
	var lastName string
	var year int
	var age int

	fmt.Println("Input your name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Println("Input your birth year:")
	fmt.Scanln(&year)

	age = 2023 - year

	fmt.Println("Your age is:", age)
	fmt.Println("=======================================")

	var fullName string
	var ageTrainer int

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Data diri trainer Enigma")
	fmt.Print("Masukan nama trainer:")

	scanner.Scan()

	fullName = scanner.Text()

	fmt.Print("Masukan umur:")

	scanner.Scan()

	ageTrainer, _ = strconv.Atoi(scanner.Text())

	fmt.Println(fullName, "|", ageTrainer)



}

