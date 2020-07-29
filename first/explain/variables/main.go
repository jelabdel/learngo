package main

import "fmt"

func main() {

	name := "Carl"
	isScientist := true
	age := 62
	degree := 5.

	// var declaration:
	// var speed int
	// var heat float64
	// var off bool
	// var brand string

	// multiple declarations:
	// Zero values
	var (
		myAge int

		heat  float64
		off   bool
		brand string
	)

	// parallel multiple variable declaration:
	var speed, velocity int

	// shorthand variable declaration and assignment:
	yourAge := 40 // var yourAge = 40

	// int
	fmt.Println(
		-200, -100, 100, 50,
	)

	// float64
	fmt.Println(
		-50.0, .4, 55.77, .0, 1.)

	// bool
	fmt.Println(
		true, false,
	)

	// string
	fmt.Println(
		"", // empty string
		"hi",
	)

	// print zero value vars
	fmt.Println(
		myAge, speed, velocity,
		yourAge, heat, off,
		isScientist, age, name, degree,
	)

	fmt.Printf("%q\n", brand)

	fmt.Println()
}
