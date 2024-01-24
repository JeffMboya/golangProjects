package main

import "fmt"

// Person struct represents a basic person with Name, Age, and Status attributes.
type Person struct {
	Name   string
	Age    int
	Status bool
}

func main() {
	// Print a message indicating the use of a struct.
	fmt.Println("This is a struct")

	// Create a Person instance named user1 with specific values.
	user1 := Person{"Jeff", 11, true}

	// Print the entire details of the user1 struct.
	fmt.Println(user1)

	// Print formatted details using %+v to include field names.
	fmt.Printf("The details are: %+v\n", user1)

	// Print formatted string with specific details from the user1 struct.
	fmt.Printf("My name is %v. I am %v years old, and I am %v\n", user1.Name, user1.Age, user1.Status)
}
