package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Aunty Donna's Coffee Cafe app")
	fmt.Println("Please rate our pizzas between 1 & 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n') // the bracket me \n is for the limit until it is supposed to read

	fmt.Println("Thanks for rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added rating to your rating", numRating+1)
	}

}
