package main

import "access-modifier/library"
import "fmt"

func main() {
	fmt.Println("Hour In A Day", library.HourInDay)
	duaHari := library.DayToHour(3)

	fmt.Println(duaHari)
}