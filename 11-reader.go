package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Rating here")

	// comma ok  or err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)

}
