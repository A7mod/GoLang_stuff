// understanding Channel Iteration & Channel Closing

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	go throwNinjaStar(channel)
	fmt.Println(<-channel)
}

func throwNinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	score := rand.Intn(10)
	channel <- fmt.Sprint("You Scored: ", score)
}
