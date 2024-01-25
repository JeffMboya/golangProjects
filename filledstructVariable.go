// Each field on the same line
package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	p1 := Person{"John", 30, "New York"}
	fmt.Println(p1)
	p2 := Person{"Alexa",
		35,
		"Goergia",
	}
	fmt.Println(p2)

	p3 := Person{Name: "Mary",
		City: "California",
	}
	fmt.Println(p3)

}
