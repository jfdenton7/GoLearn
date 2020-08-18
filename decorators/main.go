package main

import "fmt"

func main() {

	wrapper(f)

	fRef := adder()
	for i := 0; i < 5; i++ {
		fmt.Println(
			fRef(2),
		)

	}
	x := 0
	y := 0
	temp1 := func() { x = x + 1 }
	temp2 := func() { y = y + 1 }
	defer temp1()
	defer temp2()
}

func wrapper(a func()) {
	fmt.Println("before")
	a()
	fmt.Println("after")
}

func f() {
	fmt.Println("inside func")
}

// add returns a function...
func adder() func(int) int {
	// this closure inlcude a reference to sum...
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
