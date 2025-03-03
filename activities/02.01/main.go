package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	bigCount := 0
	bigWord := ""

	for key, value := range words {
		if value > bigCount {
			bigCount = value
			bigWord = key
		}
	}

	fmt.Println("Most popular word: ", bigWord)
	fmt.Println("With a count of: ", bigCount)
}
