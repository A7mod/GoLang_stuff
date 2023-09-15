package main

import "fmt"

func main() {

	// We'll discuss different fmt things here, there are a lot of things other than
	fmt.Println("Thing is I", "Fucking through some", "Print statements")
	fmt.Print("Thing is I", "Fucking through some", "Print statements") //this is like, less arranged, and no spaces etc.
	fmt.Print("\n")

	// there are some verbs? Interpolations for fmt and they work with Printf

	x := 12
	y := "Interesting Ideas"

	fmt.Printf("The type of %T .", x)    // type of data type
	fmt.Printf("The type of %d . \n", x) // integers
	fmt.Printf("There are %v \n", y)     // strings eh
	fmt.Printf("There are %T.", y)

	stockPrice := 3.50
	fmt.Printf("Each share of Gopher feed is $%.2f", stockPrice) // floats

}
