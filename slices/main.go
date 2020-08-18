package main

import (
	"fmt"
	"strings"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// take a little 'slice' out of the array of ints
	var s []int = primes[1:4]

	fmt.Println(s)

	// slices are only references TO arrays, they do not copy data
	// on assignment

	// the following is a slice LITERAL
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	// includes names at locations:
	// - names[0]
	// - names[1]
	a := names[0:2]

	b := names[1:3]

	fmt.Println(a, b)

	b[0] = "xxx"

	fmt.Println(a, b)

	fmt.Println(names)

	q := []int{2, 3, 7, 11, 13}

	fmt.Println(q)

	r := []bool{true, false, true}

	fmt.Println(r)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{4, false},
	}

	fmt.Println(s1)

	// slices also have defaults, such as the following:
	x := []int{1, 2, 3}

	x1 := x[1:]
	x2 := x[:2]
	x3 := x[:]

	fmt.Println(x1, x2, x3)

	// How do we obtain len and cap on slices?
	// s := arr[:5]
	// l = len(s)
	// c = cap(s)

	// extend length up to capacity
	// s := arr[:0]
	// s = arr[:4]

	// you can also create slices with MAKE!

	slice := make([]int, 5)
	printSlice("slice", slice)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s, len=%d, cap=%d, %v\n",
		s, len(x), cap(x), x)
}

func sliceOnSlice() {
	board := [][]string{
		{"a", "b", "c"},
		{"a", "b", "c"},
		{"1", "2", "3"},
	}

	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}
}

// how can we attempt to append Slices???
// auto allocates if capacity is too small

func appending() {
	var s []int

	// above is nil slice, append one elem
	s = append(s, 0)
	// notice how we re-assign s

	// multi-elem
	s = append(s, 1, 2, 3)

}

// much like python, we can loop over a range of elements :D

func rangeBoi() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// range gives a tuple, (index, elem)
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func supaRangeBoi() {
	pow := make([]int, 10)

	// we can just use indexes...
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// we can also just use values
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
