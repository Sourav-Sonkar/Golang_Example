package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")

	newUser := User{"Sourav", "sourav@go.dev", true, 22}
	fmt.Println(newUser)
	fmt.Printf("New User Detials is: %+v\n", newUser)
	fmt.Printf("Name: %v ,Email: %v\n", newUser.Name, newUser.Email)
	newUser.getStatus()
	newUser.changeEmail()
	//it will the change the email of newUser as it is passing a copy of newUser to change email
	fmt.Printf("Name: %v ,Email: %v\n", newUser.Name, newUser.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus(){
	fmt.Println("Is user active:",u.Status)
}

func (u User) changeEmail(){
	u.Email="new@go.dev"
	fmt.Println("Changed Email is ",u.Email)
}