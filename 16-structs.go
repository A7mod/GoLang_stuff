package main

import "fmt"

func main() {

	fmt.Println("Structs te guftagoo")

	jignesh := User{"Jignesh", "jiggubhai@selmon.com", true, 21}
	fmt.Println(jignesh)
	fmt.Printf("Jignesh details : %+v \n", jignesh)
	fmt.Printf("Name is %v and email is %v \n", jignesh.Name, jignesh.Email)
	jignesh.GetStatus()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// lets continue ahead
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)

}
