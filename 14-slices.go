package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("te welcome to slices")

	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("type %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 223
	highScores[1] = 345
	highScores[2] = 777
	highScores[3] = 153

	highScores = append(highScores, 555, 666, 991)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(highScores)
	fmt.Println(highScores) // just trying up duplicates, damn

}
