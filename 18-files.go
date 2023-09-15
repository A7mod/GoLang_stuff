package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "THis better go in a file only."

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data \n", databyte)
}
