package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web server : ")
	//PerformGetRequest()
	PerformPostJsonRequest()
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

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload?

	requestBody := strings.NewReader(`
	   {
		  "coursename": "Let's go with golang",
		  "price" : 0,
		  "platform" : "youtube.com"
       }
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data,Add("firstname", "jignesh")
	data,Add("lastname", "Mehta")
	data,Add("email", "jigneshmehta@ballam.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
}