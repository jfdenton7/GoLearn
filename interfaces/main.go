package main

import "fmt"

func main() {
	var i interface{}

	i = 42
	fmt.Printf("(%v, %T)\n", i, i)

	i = "hello"
	fmt.Printf("(%v, %T)\n", i, i)

	s, ok := i.(string)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("whoops, not a str")
	}

	i = 33
	// keeping the ok return we avoid
	// runtime err(s)
	// i.e. no PANIC occurs
	v, ok := i.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a STRING!")
	}
	typeSwitch()
}

func typeSwitch() {
	var i interface{}
	i = 33

	determineType(i)

	i = "string"

	determineType(i)
}

func determineType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("INT -> %v\n", v)
	case string:
		fmt.Printf("STR -> %v\n", v)
	default:
		fmt.Println("Neither int nor str")
	}

}
