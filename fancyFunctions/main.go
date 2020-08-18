package main

import (
	"fmt"
	"math"
)

func execFunc(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// we can declare functions inside functions!
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(execFunc(hypot))
	fmt.Println(execFunc(math.Pow))
}
