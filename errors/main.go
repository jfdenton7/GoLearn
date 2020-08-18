package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	/* format normally follows...
	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Prinf("...")
		return
	}
	fmt.Pritnln(...)
	*/
	char := 'a'
	fmt.Printf("ascii of A = %d", int(char))
}

type ErrNegativeSqrt float64

// Error, understand that if you don't
// convert to base type, it will be called again
func (ens ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(ens))
}

// when we return error and it is non nil,
// the Error() method is implicitly called!
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
