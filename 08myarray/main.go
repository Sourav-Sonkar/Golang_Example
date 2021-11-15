package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitlist [4]string

	fruitlist[0]="Apple"
	fruitlist[1]="Mango"
	fruitlist[3]="Banana"


	fmt.Println(fruitlist)
	fmt.Println(len(fruitlist))

	var vegList =[3]string{"Potato","Beans","Ladyfinger"}
	fmt.Println(vegList)
}