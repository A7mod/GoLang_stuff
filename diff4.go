package main

import "fmt"

func main() {
	fmt.Println(len("Hello World!"))
	fmt.Println("Hello World"[3]) // okkk ASCII values
	fmt.Println("hello" + "world")

	const x string = "This is gonna stay"

	x = "New areas " // here we totally get an error, cause constants bruh!!
}
