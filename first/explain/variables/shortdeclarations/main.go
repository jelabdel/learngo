package main

import "fmt"

func main() {
	name, lastName := "Nikola", "Tesla"
	fmt.Println(name, lastName)

	birth, death := 1865, 1943
	fmt.Println(birth, death)

	on, off := true, false
	fmt.Println(on, off)

	degree, ratio, heat := 10.55, 30.5, 20.
	fmt.Println(degree, ratio, heat)

	nFilese, valid, msg := 10, true, "hello"
	fmt.Println(nFilese, valid, msg)
}
