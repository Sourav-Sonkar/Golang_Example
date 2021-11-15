package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something")

	//ok comma,err,comma

	input, _ := reader.ReadString('\n')
	fmt.Println("you entered ", input)
	fmt.Printf("you entered  type %T", input)
}
