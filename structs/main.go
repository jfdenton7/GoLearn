package main

import "fmt"

type Vector struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vector{1, 2})

	v := Vector{3, 5}

	fmt.Println(v.X)

	p := &v

	// since (*p).X is cumbersome,
	// you can simply leave it out,
	// i.e.
	p.X = 300
	fmt.Println(v)

	// dealing with literals...
	// declaration block
	var (
		v1 = Vector{1, 2}
		v2 = Vector{X: 1}
		v3 = Vector{}
		p1 = &Vector{1, 2}
	)

	fmt.Printf("%v, %v, %v, %v", v1, v2, v3, p1)

}
