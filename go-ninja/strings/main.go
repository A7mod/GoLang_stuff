package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello World"

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0])
	fmt.Printf("%c", s[0])
	fmt.Println()

	fmt.Println(s[0:5])
	fmt.Println(s[:7])
	fmt.Println(s[6:])

	s += " again"
	fmt.Println(s)
	fmt.Println()

	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.HasPrefix(s, "Hello"))            // true
	fmt.Println(strings.HasSuffix(s, "Hello"))            // false
	fmt.Println(strings.Replace(s, "Hello", "Tinkle", 1)) // Tinkle World again

	s += " again World"
	fmt.Println(strings.Replace(s, "World", "Tinkle", 2)) // Hello Tinkle again again Tinkle

	// let's see the string builder

}
