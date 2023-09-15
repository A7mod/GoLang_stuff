package main

import "fmt"

// understanding returning
func items() int32 { // this int32 is for returning, the type of returning
	var x int32
	x = 1200
	return x
}

func main() {
	var vari int32
	vari = items()
	fmt.Println(vari)
}
