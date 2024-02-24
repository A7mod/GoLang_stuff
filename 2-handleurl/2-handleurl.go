package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://courses.learncodeonline.in:3000/learn/Complete-iOS-developer-swiftui"

func main() {
	fmt.Println("URL handling with Golang")
	// a new way of unfucking

	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("type? %T \n", qparams)
}

// Getting irritated by below error

// Error loading workspace: gopls was not able to find modules in your workspace. When outside of GOPATH, gopls needs to know which modules you are working on. You can fix this by opening your workspace to a folder inside a Go module, or by using a go.work file to specify multiple modules. See the documentation for more information on setting up your workspace: https://github.com/golang/tools/blob/master/gopls/doc/workspace.md.
