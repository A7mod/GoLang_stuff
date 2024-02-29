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

	smokeSignal := make(chan bool) // this is a boolean channel (true or false for when we're done)
	evilNinja := "Tommy"
	go attack(evilNinja, smokeSignal) // passed as argument? Not the right terminology
	//time.Sleep(time.Second * 2) // this is making it delay more so, we gon remove this
	fmt.Println(<-smokeSignal) // receiving the message
}

func attack(target string, attacked chan bool) { // here as well, needs to be understood
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja stars at", target)
	attacked <- true // hey the attack is done here || sendong the message
}
