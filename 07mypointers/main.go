package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of pointers")

	// var ptr *int
	// fmt.Println("the default value is ",ptr)
	myNumber := 25
	var myptr = &myNumber
	fmt.Println("The pointer address is", myptr)
	fmt.Println("The value pf pointer address is", *myptr)

	*myptr=*myptr+10
	fmt.Println("The value mynumber is ",myNumber)
}
