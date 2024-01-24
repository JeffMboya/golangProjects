package main

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float32
}

func main() {
	employee1 := Employee{} //In this case, all the fields in the struct are initialized with a default zero value of that field type.
	fmt.Println(employee1)
}
