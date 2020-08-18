package main

import (
	"fmt"
	"math"
)

// Abser will couple with any
// user defined type that
// implements all functions in the
// interface
// pointer v. value matters
// with what an interface can
// reference, which is why you
// SHOULD NOT mix
type Abser interface {
	Abs() float64
}

// interfaces
// are essentially the following construct
// (value, type)

type Vector struct {
	X, Y float64
}

type FloatBoi float64

// Abs returns the
// absolute value of f
func (f FloatBoi) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

// Magnitude couples with type via
// a receiver notation...
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// if you want to change a value a
// part of the struct, you will need a pointer...

func (v *Vector) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vector{3, 4}
	fmt.Println(v.Magnitude())

	// v implicitly given to Scale as reference (&v)
	v.Scale(3)
	fmt.Println(v.Magnitude())
	// similarly, even references can be implicitly
	// dereferenced...
	p := &v
	fmt.Println(p.Magnitude())

	var i I
	var t *T
	i = t
	i.M() // gracefully hanlde nil receiver

	i = &T{"Hello"} // M() requires T to be reference
	i.M()

	// HOWEVER, nil interfaces
	// result in runtime errors
}

// nil receivers are OK to have,
// but must be handled correctly
type T struct {
	S string
}

type I interface {
	M()
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
