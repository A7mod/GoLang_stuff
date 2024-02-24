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

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println()

	// let's see the string builder

	var sb strings.Builder
	fmt.Println("This is a string Builder:", sb.String())
	fmt.Println(sb.Cap()) // tells the capacity of the builder string
	fmt.Println(sb.Len())
	//sb.WriteString("Hein, this is a Builder?  BOB Tu?")
	sb.WriteString("Hein ")

	fmt.Println("This is a string Builder:", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	fmt.Println("------------------BREAK----------------------")
	sb.Grow(10) // this grows capacity?
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	fmt.Println("------------------BREAK----------------------")
	sb.Grow(48) // this grows capacity : but the value I put in, it's not gonna be em... Ok it changes only in 0 2 4 8 16 etc
	// so Grow() actually first doubles the current capacity, then it adds up the number which is written in the brackets. i.e 26*2+48 = 52+48 = 100  (all set!)
	fmt.Println(sb.Cap())
	fmt.Print(sb.Len())
}
