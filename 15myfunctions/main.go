package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	result := adder(3, 4)
	fmt.Println("Sum is:", result)
	proResult, somemsg := proadder(1, 5, 6, 6, 6, 8, 8)
	fmt.Println("Pro Sum is:", proResult)
	fmt.Println("Pro Message is:", somemsg)

}

func adder(x int, y int) int {
	return x + y
}

//a vardic function and also example of comma,ok function
func proadder(values ...int) (int, string) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum, "Just a string"
}

func greeter() {
	fmt.Println("Hello  from golang")
}
