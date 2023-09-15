package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano()) // this is something diff and gives a plethora of unique values.

	number := rand.Intn(100) // This gives random numbers. As we can obviously see.
	// It has something to do with seed or something.
	fmt.Println("Let's see what we have here", number) // apparantly this number shouldn't change
}
