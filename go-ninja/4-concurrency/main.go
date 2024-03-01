// understanding Channel Iteration & Channel Closing

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	//numRounds := 3
	go throwNinjaStar(channel)

	////////
	// for {
	// 	message, open := <-channel      // htis is what i shappening behind the scenes, for the for loop with range.
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }
	/////////

	for message := range channel { // after the third iteration, the for loop will be looking for more items to print in the message from the channel
		fmt.Println(message)
	}

}

func throwNinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	numRounds := 3
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You Scored: ", score)
	}
	close(channel) // htis will avoid deadlocks as the channel has been closed
}
