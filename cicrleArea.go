package main

import (
	"fmt"
	"math"
)

// Declare a constant for pi
const PI float64 = 3.142

func main() {
	//Ask the user for radius of the circle
	fmt.Print("Enter radius of the circle")
	var radius float64
	fmt.Scan(&radius)

	// Calculate area of the circle using the constant pi
	area := PI * math.Pow(radius, 2)

	//Display the results

	fmt.Printf("The area of a circle of radius %.2f is %.2f \n", radius, area)
}
