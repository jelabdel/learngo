// 	TYPE CONVERSIONS

package main

import "fmt"

func main() {
	// example 1:
	// speed := 100
	// force := 2.5

	// speed = int( float64(speed) * force )
	// fmt.Println(speed)

	// example 2:
	var apple int
	var orange int32

	apple = int(orange)

	// isDelicious := bool(orange) // can not convert!!

	orange = 65
	color := string(orange)
	fmt.Println(color)

	//to prevent unused variable error
	_ = apple

}
