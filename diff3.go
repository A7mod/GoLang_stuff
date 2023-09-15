package main

import "fmt"

func main() {
	var favFood string = "Enchiladas" // the way of string #1
	fmt.Println("My favourite food is : " + favFood)

	var handSan string
	handSan = "Must use sanitizer after such bad coding!!" // the other way of strings
	fmt.Println(handSan)

	// discussing inferring
	interestingIdea := "however is this able to detect"
	ideaSmart := 12
	ideaNumb := "And still need human assistance" // this is via the operator

	fmt.Println(interestingIdea, ideaSmart, ideaNumb)

	var intense = "This way is also feasible" // this is that other way

	fmt.Println(intense, ideaNumb)
}
