package main

import "fmt"

func main() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d milions kms\n", distance)
	fmt.Printf("Orbital period: %.3f days\n", orbital)
	fmt.Printf("Does %s has life? %t\n", planet, hasLife)

	// Argument indexing: %[2]v -> the second argument will be placed here and
	// %[1]v -> the first argument will be placed here.
	fmt.Printf(
		"%v is %v away. Think %[2]v kms! %[1]v OMG!\n",
		planet, distance,
	)
}
