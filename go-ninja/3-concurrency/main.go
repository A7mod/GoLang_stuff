// buffered channels

package main

import "fmt"

func main() {
	channel := make(chan string)
	go func() {
		channel <- "First Message"
	}()
	fmt.Println(<-channel) // Prints ---- First Message

}
