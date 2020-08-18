package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("vim-go")

	e := "hello, world"

	fmt.Println("e = ", e)

	switch_f()
	switch_time()
	defer_city()
}

func switch_f() {
	fmt.Println("GO running on...")

	switch os := runtime.GOOS; os {

	case "darwin":
		fmt.Println("OS X.")
		// implicit break here... (etc)
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switch_time() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")

	}
}

func defer_city() {
	defer goes_like_stack()
	defer fmt.Print("world!\n")

	fmt.Print("Hello, ")

}

func goes_like_stack() {
	fmt.Println("See?")

}
