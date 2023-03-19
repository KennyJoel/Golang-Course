package main

import "fmt"

func main() {
	input := "This is an example sentence."
	wordCount := 0
	letterCount := 0

	for i := 0; i < len(input); i++ {
		char := input[i]
		
		if char == ' ' || i == len(input)-1 {
			wordCount++
		}

		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			letterCount++
		}
	}

	fmt.Printf("Input: %s\nWord count: %d\nLetter count: %d\n", input, wordCount, letterCount)
}
