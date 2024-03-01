// understanding Channel Iteration & Channel Closing

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	numRounds := 3
	go throwNinjaStar(channel, numRounds)
	for i := 0; i < numRounds; i++ {
		fmt.Println(<-channel)
	}

}

func throwNinjaStar(channel chan string, numRounds int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You Scored: ", score)
	}
}
