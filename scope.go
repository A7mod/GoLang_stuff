package main

import "fmt"

var x string = "This is defined globally !"
var y string = "How will I do this now?"

func main() {

	fmt.Println(x)

}

func item() {
	fmt.Println(y) // why is this not printing man
}
