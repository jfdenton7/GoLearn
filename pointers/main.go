package main

import "fmt"

func main() {
	i, j := 32, 33

	p := &i // (p points to i)

	fmt.Println(*p) // dereference operator the same as "C"

	// we can change the value i holds by changing p, since p points to i
	*p = 50
	fmt.Println(i)

	p = &j
	*p = *p / 37 // dereference operator takes precedence over '/'
	fmt.Println(j)

}
