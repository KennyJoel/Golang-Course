package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numProblems := 5 // change the number of problems here
	fmt.Printf("I will solve %d %s:\n", numProblems)
	for i := 1; i <= numProblems; i++ {
		fmt.Printf("Problem %d: %d + %d = _____\n", i, rand.Intn(10), rand.Intn(10))
	}
}