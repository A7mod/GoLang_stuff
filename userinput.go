package main

import "fmt"

func main() {
	fmt.Println("How can I be of service?")

	var serve string
	fmt.Scan(&serve) // taking user input using Scan() ---- "&" is for address

	// Although it can only take one argument (string) at a time.
	// for maybe 2? here is the way.

	var vice string
	fmt.Scan(&vice)

	fmt.Printf("Can you fetch %v %v bottles from the fridge?", vice, serve) // had to call these aage peeche since the reading is done in a different way.

}
