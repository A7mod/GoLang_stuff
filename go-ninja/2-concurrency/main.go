package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}() //this is a function call, that's how this defer func works

	evilNinja := "Tommy"
	go attack(evilNinja)
	time.Sleep(time.Second * 2)
}

func attack(target string) {
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja stars at", target)
}
