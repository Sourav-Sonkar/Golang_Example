package main

import "fmt"

//public variable ie start with capital letter
const LoginToken string = "dfgfgf"

func main() {
	//fmt.Println("Variables")
	var username string = "Sourav"
	fmt.Println(username)
	fmt.Printf("Variables is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type: %T \n", smallVal)

	var floatVal float64 = 255.5555555555555555555555555555555
	fmt.Println(floatVal)
	fmt.Printf("Variables is of type: %T \n", floatVal)

	//default values and aliases
	var anothervarible string
	fmt.Println(anothervarible)
	fmt.Printf("Variables is of type: %T \n", anothervarible)

	// implicit types
	var website = "dummy.com"
	fmt.Println(website)

	//no var style
	numberofUsers := 5000
	fmt.Println(numberofUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variables is of type: %T \n", LoginToken)
}
