package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["C"] = "C/C++"
	languages["PY"] = "Python"

	fmt.Println("List of all langs: ", languages)

	fmt.Println("C is short for: ", languages["C"])

	delete(languages, "C") // delete aise karein

	fmt.Println("list after laat maari: ", languages)

}
