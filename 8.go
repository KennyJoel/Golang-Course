package main

import "fmt"

const Pi = 69

func main() {
	const World = "Kenny"
	fmt.Println("Welcome", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = "false?"
	fmt.Println("Is it true?", Truth)
}