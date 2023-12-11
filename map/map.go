package main

import "fmt"

func main() {
	user := map[string]string{
		"username": "adrian",
		"email": "adrianhhd@email.com",
	}

	fmt.Println(user)

	fmt.Println("===============================")

	var scores = make(map[string]int)
	fmt.Println(scores)

	scores["Java"] = 85
	scores["React"] = 95

	fmt.Println(scores)
	fmt.Println("Nilai Java:", scores["Java"])

	fmt.Println("===============================")

	delete(scores, "Java")
	fmt.Println(scores)

	fmt.Println("===============================")

	names := map[int]string {
		1: "Adrian",
		2: "Hadinata",
		3: "Hadi",
		4: "Darsono",
	}

	for key, value := range names {
		fmt.Println(key, value)
	}
}