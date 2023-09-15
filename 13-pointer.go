package main

import "fmt"

func main() {

	fmt.Println("Te a class on pointersch: ")

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of actual pointer : ", ptr)
	fmt.Println("Value of actual pointer : ", *ptr)
}
