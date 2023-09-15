package main

import "fmt"

// I am not able to import "fmt"

func namesake() {

	fmt.Println("I am a friend of Nicole!")
	// could also understand scope here
	x := 12
	fmt.Println(x)

}

func main() {
	namesake()
	//fmt.Println(x) // since this will not print here : basic scope rules

	namesake()

}
