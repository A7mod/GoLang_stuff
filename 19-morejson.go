package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("JSON items")
	EncodeJson()
}

func EncodeJson() {
	devCourses := []course{
		{"ReactJS Bootcamp", 299, "Learnonline.in", "abc123", []string{"web-dev", "js"}},
		{"Python Bootcamp", 399, "Learnonline.in", "abc223", []string{"full-stack", "js"}},
		{"Node Bootcamp", 1299, "Learnonline.in", "abc323", []string{"google", "js"}},
		{"CP Bootcamp", 1199, "Learnonline.in", "abc423", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(devCourses, "", "\t") //cool stuff
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)

}
