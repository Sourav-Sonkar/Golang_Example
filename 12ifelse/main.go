package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	var loginCount int = 10
	var result string

	if loginCount > 10 {
		result = "Regular user"
	} else if loginCount < 10 {
		result = "qwerty"
	} else {
		result = "something  else"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is smaller than 10")
	} else {
		fmt.Println("Number is large or equal to 10")
	}

	// if err != nil {
	// }

}
