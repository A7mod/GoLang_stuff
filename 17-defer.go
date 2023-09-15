package main

import "fmt"

func main() {

	defer fmt.Println("World")
	defer fmt.Println("One") // so what's happening here. the order is reverse if the defers are in order
	defer fmt.Println("Two")
	fmt.Println("hello")
}
