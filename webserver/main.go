package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to web server : ")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code :", response.StatusCode)
	fmt.Println("COntent length :", response.ContentLength)

	content, _ := iotuil.ReadAll(response.Body)

	fmt.Println(string(content))

}
