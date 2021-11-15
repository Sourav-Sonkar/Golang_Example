package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")

	newUser := User{"Sourav", "sourav@go.dev", true, 22}
	fmt.Println(newUser)
	fmt.Printf("New User Detials is: %+v\n", newUser)
	fmt.Printf("Name: %v ,Email: %v\n", newUser.Name, newUser.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
