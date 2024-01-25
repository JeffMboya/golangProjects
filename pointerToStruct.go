package main

import "fmt"

// employee is a struct representing an employee with name, age, and salary.
type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	// Create a new employee using the 'new' function, which returns a pointer to an initialized zero-value employee.
	empP := new(employee)

	// Print the memory address of the employee pointer.
	fmt.Printf("Emp Pointer Address: %p\n", empP)

	// Print the details of the employee pointer (pointer value).
	fmt.Printf("Emp Pointer: %+v\n", empP)

	// Dereference the employee pointer to access the actual employee value and print its details.
	fmt.Printf("Emp Value: %+v\n", *empP)
}
