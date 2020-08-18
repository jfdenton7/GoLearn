package main

import "fmt"

// UpperCamelCase is considered public

// Person is just a boi/girly
type Person struct {
	Name string
	Age  int
}

// String is implementing the stringer
// interface, which is called in printf
func (p Person) String() string {
	return fmt.Sprintf("name: %v, age: %v", p.Name, p.Age)
}

func main() {

	a := Person{"Arthur", 17}
	b := Person{"James", 32}

	fmt.Println(a, b)
}
