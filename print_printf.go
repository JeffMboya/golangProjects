
/*The print function is a basic function that simply concatenates its arguments and writes the result to the standard output.
It does not provide any formatting options. It adds spaces between the arguments.*/


package main

import "fmt"

func main() {
    fmt.Print("Hello", "World", 42)
}


/*The printf function is a formatted printing function. It allows you to specify a format string that includes placeholders for values, 
and it replaces these placeholders with the corresponding values. It provides more control over the output's format and alignment. 
It does not automatically add spaces between the arguments; the formatting is controlled by the format string.*/

package main

import "fmt"

func main() {
    name := "John"
    age := 25
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
