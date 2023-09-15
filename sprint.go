package main

import "fmt"

func main() {
	grade := "100"
	compliment := "Great job!"
	teacherSays := fmt.Sprint("You scored a ", grade, " on the test! \n", compliment)

	fmt.Print(teacherSays)

	step1 := "Breathe in..."
	step2 := "Breathe out... \n"

	meditation := fmt.Sprintln(step1, step2)

	fmt.Println(meditation)

	// we can use Sprintf like this :

	item := "She got em big %v."

	ding := "package"

	var work string

	work = fmt.Sprintf(item, ding)

	fmt.Print(work) // well that's a lot of work
}
