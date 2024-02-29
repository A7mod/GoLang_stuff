// buffered channels

package main

import "fmt"

func main() {
	channel := make(chan string)

	// trying to understand if this works this way. This is step 2. It will fail (deadlock)
	channel <- "First Message"
	// go func() {
	// 	channel <- "First Message"
	// }()
	fmt.Println(<-channel) // Prints ---- First Message

	//initially the main go routine is running, then the other go routine starts when
	// we use the go func(). (The initial go routine is continuing while all this is happening
	// at some point we are sending a message to the channel.
	// apparently the channel has a default capacity of 0
	// so when we send the "message" from the task go routine to the main go routin, to receive the message
	// otherwise it will cause a congestion, since the channel has 0 capacity

}
