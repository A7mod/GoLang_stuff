package main

import "fmt"

func mul(x, y int32) int32 { // that's how we define parameters in functions
	return x * y // could also be done like x int32, y int32
}

func main() {
	var product int32
	product = mul(22, 40)
	fmt.Println(product)
}
