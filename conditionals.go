// Most of the conditionals are same only, but let's give some examples.
package main

import "fmt"

func main() {
	night := "Nacht"
	if night == "Abend" {
		fmt.Println("Yes yes")

	} else { // apparantly Go is a very anal language about syntax. The else statement has to be just after the if block, not a new line.
		fmt.Println("Nein Nein")
	}

	// scoped short declaration statement.

	if success := false; success { // and this is the way this is done, similarly for switch case statements, tho not doing it.. .. .. . .. ... ...
		fmt.Println("Yes yes")
	} else {
		fmt.Println("This ain't working mate!")
	}

}
