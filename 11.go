package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numProblems := 5 // change the number of problems here
	fmt.Printf("I will solve %d %s:\n", numProblems, pluralize(numProblems, "problem", "problems"))
	for i := 1; i <= numProblems; i++ {
		fmt.Printf("Problem %d: %d + %d = _____\n", i, rand.Intn(10), rand.Intn(10))
	}
}

func pluralize(num int, singular string, plural string) string {
	if num == 1 {
		return singular
	}
	return plural
}
