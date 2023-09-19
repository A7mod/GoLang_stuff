package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json: "website"`
	Password string   `json:"-"`              // this is important for passwords and secrets
	Tags     []string `json:"tags,omitempty"` // this is a nice thing
}

func main() {
	fmt.Println("JSON items")
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"Platform": "Learnonline.in",
		"tags": ["web-dev","js"]
    }
	`)

	var devCourses course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &devCourses)
		fmt.Printf("%#v \n", devCourses)

	} else {
		fmt.Println("JSON WAS NOT VALID MATE")
	}

	// some cases where you just want to add data to key value...

	var myOnlineData map[string]interface{} //interface type is defined as a set of method signatures..
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for z, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v, und Type is : %T \n", z, v, v)
	}

}
