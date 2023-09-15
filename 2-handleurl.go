package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://courses.learncodeonline.in/learn/Complete-iOS-developer-swiftui"

func main() {
	fmt.Println("URL handling with Golang")
	// a new way of unfucking

	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	// qparams := result.Query()
	// fmt.Printf("type? %T \n", qparams)
}
