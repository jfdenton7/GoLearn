package main

import (
	"fmt"
	"strings"
)

type Point struct {
	Lat, Long float64
}

func main() {
	var m map[string]Point
	// string is key, Point is value...
	m = make(map[string]Point)

	m["Bell Labs"] = Point{
		40.68433, -74.39967,
	}

	// access to elem by key:"Bell Labs"
	fmt.Println(m["Bell Labs"])

	// we can also declare map literals...

	myMap := map[string]Point{
		"Bell Labs": {
			40.7, -74.4, // no need to specify type, redundant
		},
		"Google": {
			37.4, -122.1,
		},
	}
	fmt.Println(myMap)

	mutableMaps()

	m1 := WordCount("hello hello test 123 123 123")
	fmt.Println("And the answer is .... \n\n", m1)
}

func mutableMaps() {
	m := make(map[string]int)

	m["Answer"] = 42

	fmt.Println("the value is:", m["Answer"])

	// now we change the value that the key points to
	m["Answer"] = 38
	fmt.Println("the value is:", m["Answer"])

	// to delete a key:value pair, we perforn
	delete(m, "Answer")
	fmt.Println("the value is:", m["Answer"])

	v, ok := m["Answer"]

	if !ok {
		fmt.Println("whoops, \"Answer\" is not in the map")
	} else {
		fmt.Println("the value is:", v)
	}
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		// if val found in map,
		// we want to add to the previous val
		// OR we will set
		if _, ok := m[word]; ok {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}

	return m
}
