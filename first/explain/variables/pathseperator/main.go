package main

import (
	"fmt"
	"path"
)

func main() {
	// var dir, file string
	// dir, file = path.Split("css/main.css")
	// fmt.Println("dir :", dir)
	// fmt.Println("file :", file)

	// using short declaration:
	// dir, file := path.Split("css/main.css")
	// fmt.Println("file :", file)
	// fmt.Println("dir :", dir)

	// if you only want to extract the file name:
	_, file := path.Split("css/main.css")
	fmt.Println("file :", file)
}
