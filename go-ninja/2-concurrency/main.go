package main

import (
	"fmt"
	"time"
)

func main() {

	evilNinja := "Tommy"
	go attack(evilNinja)
	time.Sleep(time.Second * 2)
}

func attack(target string) {
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja stars at", target)
}
