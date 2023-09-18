package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	// fmt.Println(string(content)) // way 1 to do it

	fmt.Println("Bytecount is :", byteCount)
	fmt.Println(responseString.String()) // way 2 by using builder from Strings

}
