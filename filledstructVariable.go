package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	//Each field on the same line
	p1 := Person{"John", 30, "New York"}
	fmt.Println(p1)

	//Each field on a different line
	p2 := Person{"John",
		30,
		"New York",
	}
	fmt.Println(p2)

}
