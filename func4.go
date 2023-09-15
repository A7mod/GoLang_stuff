package main

import "fmt"

// this is pretty simple, defer just delays the "delay item" until the end of the function or whenever it is called. Samjha?

func minute(nine, eight int32) int32 {
	fmt.Println("Just let it happen man")
	defer fmt.Println("Noone remembers vaade!")

	fmt.Println("Just let it  man")
	fmt.Println("Just let it happen ")
	return nine * eight

}

func main() {
	fmt.Println(minute(12, 22))
}
