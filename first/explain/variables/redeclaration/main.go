package main

import "fmt"

// REDECLARATION possible only when new variable is used on left side
func main() {
	name := "Nikola"

	name, age := "Marie", 66 // age is a new variable
	fmt.Println(name, age)

	// name = "Abert" // assignment
	// birth := 1879  // declaration and initialization

	// ^--- same as above
	name, birth := "Albert", 1879 // redacleration (birth is a new var)
	fmt.Println(name, birth)

}
