package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitlist = []string{"Peach"}
	fmt.Printf("This is type %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Apple", "banana", "Tomato")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highScores := make([]int, 4)
	highScores[0] = 4
	highScores[1] = 55
	highScores[2] = 66
	highScores[3] = 888
	//highScores[4]=888 will not work

	highScores = append(highScores, 444, 666, 77)
	fmt.Println(sort.IntsAreSorted(highScores))

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to reomove a values from slices based on index

	var alphabet=[]string{"A","B","C","D","E"}
	fmt.Println(alphabet)
	var index int=2
	alphabet=append(alphabet[:index],alphabet[index+1:]...)
	fmt.Println(alphabet)

}
