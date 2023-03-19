package main

import (
	"fmt"
	"math"
)

func main() {
	var g, f int = 5, 6
	var k float64 = math.Sqrt(float64(g*f + g*g))
	var h uint = uint(k)
	fmt.Println(g, k, h)
}
